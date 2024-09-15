package logic

import (
	"context"
	"errors"
	"server/common/list_query"
	common "server/common/models"
	chat "server/models/chat"
	user "server/models/user"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"server/chat/api/internal/svc"
	"server/chat/api/internal/types"
	"server/user/rpc/proto"
)

type ChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryLogic {
	return &ChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ChatUserInfo struct {
	ID       uint   `json:"id"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
}

type ChatHistory struct {
	ID        uint              `json:"id"`
	SendUser  ChatUserInfo      `json:"sendUser"`
	RevUser   ChatUserInfo      `json:"revUser"`
	IsMe      bool              `json:"isMe"`       // 哪条消息是我发的
	CreatedAt string            `json:"created_at"` // 消息时间
	Msg       *common.Message   `json:"msg"`
	SystemMsg *common.SystemMsg `json:"systemMsg"`
	ShowDate  bool              `json:"showDate"` // 是否显示时间
}
type ChatHistoryResponse struct {
	List  []ChatHistory `json:"list"`
	Count int64         `json:"count"`
}

func (l *ChatHistoryLogic) ChatHistory(req *types.ChatHistoryRequest) (resp *ChatHistoryResponse, err error) {
	var friendShip user.FriendshipModel
	if !friendShip.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("你们还不是好友呢")
	}
	// 是好友的话查询聊天记录
	chatList, count, err := list_query.ListQuery(l.svcCtx.DB, chat.ChatModel{}, list_query.Option{
		PageInfo: common.PageInfo{
			Page:  req.Page,
			Limit: req.PageSize,
		},
		Sort: "created_at desc",
		Where: l.svcCtx.DB.Where(`((send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)) and 
			id not in (select chat_id from user_chat_delete_models where user_id = ?)
			`,
			req.UserID, req.FriendID, req.FriendID, req.UserID, req.UserID, // 不查询已经删除的聊天记录
		),
	})
	if err != nil {
		logx.Errorf("查询聊天记录失败: %v", err)
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("你们还没有聊天记录呢")
	}

	// 获取到了所有的聊天记录后, 要将对应的聊天记录进行归类，按照不同的用户进行归类
	// 消息是按时间降序得到的，所以idx在前面的是最新的消息
	userIDs := []uint32{uint32(req.UserID), uint32(req.FriendID)}
	usersInfo, err := l.svcCtx.UserRpc.UserListInfo(context.Background(), &proto.UserListInfoRequest{UserIDs: userIDs})
	if err != nil {
		return nil, err
	}

	chatHistoryList := make([]ChatHistory, 0)

	// 用于记录当前页的聊天记录是否需要显示时间, 因为消息是按照时间降序排列的，列表最后一条记录是最早的记录
	// 判断这最早的那条记录是否超过一天时间
	isShowDate := false
	if chatList[len(chatList)-1].CreatedAt.Before(time.Now().Add(-time.Hour * 24)) {
		isShowDate = true
	}

	for idx, chat := range chatList {
		// 判断当前页的聊天记录是否需要显示时间
		chatHistory := ChatHistory{
			ID:        chat.ID,
			CreatedAt: chat.CreatedAt.Format("2006-01-02 15:04:05"),
			Msg:       chat.Msg,
			SystemMsg: chat.SystemMsg,
			ShowDate:  false,
			SendUser: ChatUserInfo{
				ID:       chat.SendUserID,
				NickName: usersInfo.UserInfoMap[uint32(chat.SendUserID)].NickName,
				Avatar:   usersInfo.UserInfoMap[uint32(chat.SendUserID)].Avatar,
			},
			RevUser: ChatUserInfo{
				ID:       chat.RevUserID,
				NickName: usersInfo.UserInfoMap[uint32(chat.RevUserID)].NickName,
				Avatar:   usersInfo.UserInfoMap[uint32(chat.RevUserID)].Avatar,
			},
			IsMe: chat.SendUserID == req.UserID,
		}
		if idx == len(chatList)-1 {
			chatHistory.ShowDate = isShowDate
		}
		chatHistoryList = append(chatHistoryList, chatHistory)

	}
	resp = &ChatHistoryResponse{
		List:  chatHistoryList,
		Count: count,
	}
	return
}
