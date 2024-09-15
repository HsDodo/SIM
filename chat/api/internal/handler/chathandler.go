package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
	"net/http"
	"server/chat/api/internal/logic"
	"server/chat/api/internal/svc"
	"server/chat/api/internal/types"
	common "server/common/models"
	"server/common/response"
	chat "server/models/chat"
	user "server/models/user"
	user_proto "server/user/rpc/proto"
	"time"
)

type UserWsInfo struct {
	UserInfo    user.UserModel             // 用户信息
	WsClientMap map[string]*websocket.Conn // 保存用户的websocket连接
	CurrentConn *websocket.Conn            // 当前用户的websocket连接
}

var OnlineUserWsMap = map[uint]*UserWsInfo{}

func chatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("参数解析失败: %s", err.Error())
			response.Response(w, nil, err)
			return
		}
		upGrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		// 将http链接升级为websocket链接
		conn, err := upGrader.Upgrade(w, r, nil)
		if err != nil {
			logx.Errorf("websocket连接失败: %s", err.Error())
			response.Response(w, nil, err)
			return
		}

		addr := conn.RemoteAddr().String() // 获取发起方的地址
		// 这里 defer 关闭连接，并且将该客户端信息进行更新
		defer func() {
			conn.Close()
			userWsInfo, ok := OnlineUserWsMap[req.UserID]
			if ok {
				// 删除这条连接信息
				delete(userWsInfo.WsClientMap, addr)
			}
			if userWsInfo != nil && len(userWsInfo.WsClientMap) == 0 {
				// 如果该用户的所有连接都已经关闭，那么删除这个用户的信息
				delete(OnlineUserWsMap, req.UserID)
				// 在redis将用户的在线状态删除
				svcCtx.Redis.HDel(context.Background(), "online", fmt.Sprintf("%d", req.UserID))
				svcCtx.Redis.HDel(context.Background(), "ws_online_users", fmt.Sprintf("%d", req.UserID))
			}
		}()

		// --------------------------------------------------------------
		//Todo: 调用用户rpc服务,获取用户信息, 并处理用户登录
		res, err := svcCtx.UserRpc.UserInfo(context.Background(), &user_proto.UserInfoRequest{
			UserId: uint32(req.UserID),
		})
		if err != nil {
			logx.Errorf("用户信息获取失败: %s", err.Error())
			response.Response(w, nil, err)
			return
		}

		var userInfo user.UserModel
		err = json.Unmarshal(res.Data, &userInfo)
		if err != nil {
			logx.Errorf("用户信息解析失败: %s", err.Error())
			response.Response(w, nil, err)
			return
		}

		isFirstLogin := false
		// 用户上线了，需要将其信息保存
		userWsInfo, ok := OnlineUserWsMap[req.UserID]
		if !ok { // 用户第一次上线,
			OnlineUserWsMap[req.UserID] = &UserWsInfo{
				UserInfo:    userInfo,
				WsClientMap: map[string]*websocket.Conn{addr: conn},
				CurrentConn: conn,
			}
			isFirstLogin = true
			// 在redis中保存用户的在线状态
			svcCtx.Redis.HSet(context.Background(), "online", fmt.Sprintf("%d", req.UserID), "1")
			svcCtx.Redis.HSet(context.Background(), "ws_online_users", fmt.Sprintf("%d", req.UserID), "1")
		} else {
			// 换设备登录了, 不需要更新redis中的在线状态
			userWsInfo.WsClientMap[addr] = conn // 添加新设备
			userWsInfo.CurrentConn = conn       // 更新当前设备
		}

		// ----------------------------------------------------------------------------------------------
		//Todo:给好友发送上线提醒，前提是好友在线并且设置了在线提醒
		friendListResp, err := svcCtx.UserRpc.FriendList(context.Background(), &user_proto.FriendListRequest{
			UserId: uint32(req.UserID),
		})
		friendsInfoMap := friendListResp.UserInfoMap
		if isFirstLogin { // 第一次登录，需要给好友发送上线提醒
			for _, friendInfo := range friendsInfoMap {
				// 先判断好友在redis中的在线状态
				onlineStatus, err := svcCtx.Redis.HGet(context.Background(), "ws_online_users", fmt.Sprintf("%d", friendInfo.UserId)).Result()
				if err != nil {
					//logx.Error(err)
					logx.Infof("该好友不在线: %s", friendInfo.NickName)
					continue
				}
				friendWsInfo, ok := OnlineUserWsMap[uint(friendInfo.UserId)]
				if ok && onlineStatus == "1" {
					// 判断是否开启了上线提醒
					if friendWsInfo.UserInfo.UserConf.FriendOnline {
						textMsg := fmt.Sprintf("%s上线了", userInfo.Nickname)
						loginNotifyMsg := common.Message{
							Type:        common.FriendOnlineMsgType,
							TextMessage: &textMsg,
						}
						// Todo: 消息入库（待实现）
						msgID, err := StoreMsg(svcCtx.DB, req.UserID, uint(friendInfo.UserId), &loginNotifyMsg)
						if err != nil {
							logx.Error(err)
							continue
						}
						// 发送消息
						err = SendMsg(req.UserID, uint(friendInfo.UserId), &loginNotifyMsg, msgID)
						if err != nil {
							logx.Error(err)
							continue
						}
					}
				}
			}
		}
		// ----------------------------------------------------------------------------------------------
		// Todo: 发送消息, 不同消息不同处理
		for {
			// 从客户端读取消息
			_, msgData, err := conn.ReadMessage() // 客户端发送 ws 请求， 参数为 common.Message 类型的json字符串
			if err != nil {
				logx.Info("用户断开连接,没有取到值！")
				logx.Error(err)
				break
			}
			// 判断用户是否被禁言
			if userInfo.UserConf.ForbidChat {
				err = SendTipMsg(conn, "您已被禁止聊天!")
				continue
			}
			logx.Infof("收到消息: %s", string(msgData))
			chatMsgReq := ChatMsgRequest{}
			err = json.Unmarshal(msgData, &chatMsgReq)
			if err != nil {
				logx.Error(err)
				SendTipMsg(conn, "消息解析失败")
				continue
			}

			logx.Infof("收到消息ChatMsgReq: %s", chatMsgReq)

			// 判断是否是好友
			friendShip, err := svcCtx.UserRpc.IsFriend(context.Background(), &user_proto.IsFriendRequest{
				UserId:   uint32(req.UserID),
				FriendId: uint32(chatMsgReq.RecvUserID),
			})
			if err != nil {
				logx.Error(err)
				SendTipMsg(conn, "好友关系查询失败!")
				continue
			}

			if !friendShip.IsFriend {
				err = SendTipMsg(conn, "对方不是您的好友!")
				continue
			}
			// 判断消息类型
			if !(chatMsgReq.Msg.Type >= 1 && chatMsgReq.Msg.Type <= 14) {
				err = SendTipMsg(conn, "消息类型错误!")
				continue
			}
			// 判断消息是否格式正确
			if err = chatMsgReq.Msg.IsValid(); err != nil {
				SendTipMsg(conn, err.Error())
				continue
			}
			// 判断消息类型
			switch chatMsgReq.Msg.Type {
			case common.TextMsgType:

			case common.FileMsgType:
				err = handleFileMsg(chatMsgReq, svcCtx, conn, req.UserID)
				if err != nil {
					logx.Error(err)
				}
				continue
			case common.WithdrawMsgType:
				err = handleWithdrawMsg(svcCtx, chatMsgReq, conn, OnlineUserWsMap, req.UserID)
				if err != nil {
					logx.Error(err)
				}
				continue
			case common.ReplyMsgType:
				err = handleReplyMsg(svcCtx, chatMsgReq, conn, OnlineUserWsMap, req.UserID)
				if err != nil {
					logx.Error(err)
				}
				continue
			case common.VideoMsgType:
				err = handleVideoCallMsg(svcCtx, chatMsgReq, conn, OnlineUserWsMap, req.UserID)
				if err != nil {
					logx.Error(err)
				}
			}

			// 将消息入库并发送给好友
			msgID, err := StoreMsg(svcCtx.DB, req.UserID, chatMsgReq.RecvUserID, chatMsgReq.Msg)
			if err != nil {
				logx.Error(err)
				SendTipMsg(conn, "消息入库失败")
				continue
			}
			SendMsg(req.UserID, chatMsgReq.RecvUserID, chatMsgReq.Msg, msgID)

		}

		l := logic.NewChatLogic(r.Context(), svcCtx)
		resp, err := l.Chat(&req)
		response.Response(w, resp, err)

	}
}

type ChatMsgRequest struct {
	RecvUserID uint            `json:"recvUserID"`
	Msg        *common.Message `json:"msg"`
}

type ChatMsgResponse struct {
	ID         uint            `json:"id"`
	IsMe       bool            `json:"isMe"`
	RecvUser   UserInfo        `json:"recvUser"`
	SendUser   UserInfo        `json:"sendUser"`
	Msg        *common.Message `json:"msg"`
	CreatedAt  time.Time       `json:"createdAt"`
	MsgPreview string          `json:"msgPreview"`
}

type UserInfo struct {
	UserID   uint   `json:"userID"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

func SendTipMsg(conn *websocket.Conn, msg string) error {
	tipMsg := common.Message{
		Type: common.TipMsgType,
		TipMessage: &common.TipMessage{
			Status:  "error",
			Content: msg,
		},
	}
	tipMsgData, err := json.Marshal(tipMsg)
	if err != nil {
		logx.Error(err)
		return errors.Errorf("消息解析失败")
	}
	err = conn.WriteMessage(websocket.TextMessage, tipMsgData)
	if err != nil {
		logx.Error(err)
		return errors.Errorf("消息发送失败")
	}
	return nil
}

// StoreMsg Todo: 消息入库
func StoreMsg(db *gorm.DB, senderID uint, recvID uint, msg *common.Message) (msgID uint, err error) {

	if msg == nil {
		return 0, errors.New("消息不能为空")
	}

	newChat := chat.ChatModel{
		SendUserID: senderID,
		RevUserID:  recvID,
		MsgType:    msg.Type,
		Msg:        msg,
		MsgPreview: msg.MsgPreview(),
	}

	err = db.Create(&newChat).Error
	if err != nil {
		logx.Error(err)
		return
	}
	return newChat.ID, nil
}

func SendMsg(senderID uint, recverID uint, msg *common.Message, msgID uint) error {
	// 获取通信双方的websocket连接
	senderWsInfo, ok1 := OnlineUserWsMap[senderID]
	recverWsInfo, ok2 := OnlineUserWsMap[recverID]
	if !ok1 || !ok2 {
		logx.Error("用户不在线")
		return errors.Errorf("用户不在线")
	}

	chatMsgResp := ChatMsgResponse{
		ID: msgID,
		RecvUser: UserInfo{
			UserID:   recverWsInfo.UserInfo.ID,
			NickName: recverWsInfo.UserInfo.Nickname,
			Avatar:   recverWsInfo.UserInfo.Avatar,
		},
		SendUser: UserInfo{
			UserID:   senderWsInfo.UserInfo.ID,
			NickName: senderWsInfo.UserInfo.Nickname,
			Avatar:   senderWsInfo.UserInfo.Avatar,
		},
		IsMe:       senderID == recverID,
		Msg:        msg,
		CreatedAt:  time.Now(),
		MsgPreview: msg.MsgPreview(),
	}

	chatMsgRespData, err := json.Marshal(chatMsgResp)
	if err != nil {
		logx.Error(err)
		return errors.Errorf("消息解析失败")
	}
	// 将消息发送给接收方的所有设备
	err = SendMsgToAllWsClients(chatMsgRespData, recverWsInfo.WsClientMap)
	if err != nil {
		logx.Error(err)
		return err
	}
	return nil
}

func SendMsgToAllWsClients(msg []byte, wsClients map[string]*websocket.Conn) error {
	for _, wsConn := range wsClients {
		err := wsConn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			logx.Error(err)
			return errors.Errorf("消息发送失败")
		}
	}
	return nil
}

func SendMsgWithoutStore(senderID uint, recverID uint, msg *common.Message) error {
	// 获取通信双方的websocket连接
	senderWsInfo, ok1 := OnlineUserWsMap[senderID]
	recverWsInfo, ok2 := OnlineUserWsMap[recverID]
	if !ok1 || !ok2 {
		logx.Error("用户不在线")
		return errors.Errorf("用户不在线")
	}

	chatMsgResp := ChatMsgResponse{
		RecvUser: UserInfo{
			UserID:   recverWsInfo.UserInfo.ID,
			NickName: recverWsInfo.UserInfo.Nickname,
			Avatar:   recverWsInfo.UserInfo.Avatar,
		},
		SendUser: UserInfo{
			UserID:   senderWsInfo.UserInfo.ID,
			NickName: senderWsInfo.UserInfo.Nickname,
			Avatar:   senderWsInfo.UserInfo.Avatar,
		},
		IsMe:       senderID == recverID,
		Msg:        msg,
		CreatedAt:  time.Now(),
		MsgPreview: msg.MsgPreview(),
	}

	chatMsgRespData, err := json.Marshal(chatMsgResp)
	if err != nil {
		logx.Error(err)
		return errors.Errorf("消息解析失败")
	}
	// 将消息发送给接收方的所有设备
	err = SendMsgToAllWsClients(chatMsgRespData, recverWsInfo.WsClientMap)
	if err != nil {
		logx.Error(err)
		return err
	}
	return nil
}
