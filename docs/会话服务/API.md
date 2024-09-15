### 会话服务 API

#### 主要功能

> 1. 查询用户聊天记录
> 2. 查询最近会话列表

#### API配置文件

```yaml
Host: 0.0.0.0
Port: 9003
Name: "chat.api"

Etcd:
  Hosts:
    - "localhost:2379"

# 用户微服务
UserRpc:
  Etcd:
    Hosts:
      - "localhost:2379"
    Key: user.rpc

Mysql:
  DataSource: "root:root@tcp(localhost:3306)/sim_db?charset=utf8&parseTime=True&loc=Local"

Log:
  ServiceName: auth
  Encoding: plain
  Stat: false
  TimeFormat: 2006-01-02 15:04:05

Redis:
  Addr: "127.0.0.1:6379"
  Pwd: "5566"
  DB: 0

```

#### 加载ServiceContext

```go
package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"server/chat/api/internal/config"
	"server/common/Interceptor"
	"server/common/logger"
	"server/core"
	user_rpc "server/user/rpc/userservice"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.Mysql.DataSource
	db := core.InitGorm(dsn)
	redisClient, err := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	if err != nil {
		logger.Fatalf("redis初始化失败!: %v", err)
	}
	return &ServiceContext{
		Config:  c,
		DB:      db,
		Redis:   redisClient,
		UserRpc: user_rpc.NewUserService(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(Interceptor.ClientInfoInterceptor))),
	}
}

```

#### API文件

```go
syntax = "v1"

type ChatHistoryRequest {
	UserID   uint `header:"userID"`
	Page     int  `form:"page,optional"`
	PageSize int  `form:"pageSize,optional"`
	FriendID uint `form:"friendID"`
}

type ChatHistoryResponse {}

type ChatSessionRequest {
	UserID   uint `header:"userID"`
	Page     int  `form:"page,optional"`
	PageSize int  `form:"pageSize,optional"`
	Key      int  `form:"key,optional"`
}

type ChatSession { //就是界面那个最近会话列表
	UserID     uint   `json:"userID"`
	Avatar     string `json:"avatar"`
	Nickname   string `json:"nickname"`
	CreatedAt  string `json:"created_at"` // 消息时间
	MsgPreview string `json:"msgPreview"` // 消息预览
	IsTop      bool   `json:"isTop"` // 是否置顶
	IsOnline   bool   `json:"isOnline"` // 是否在线
}

type ChatSessionResponse {
	List  []ChatSession `json:"list"`
	Count int           `json:"count"`
}

service chat {
	@handler chatHistory
	get /api/chat/history (ChatHistoryRequest) returns (ChatHistoryResponse) // 查询用户聊天记录

	@handler chatSession
	get /api/chat/session (ChatSessionRequest) returns (ChatSessionResponse) //最近会话列表
//    @handler userTop
//    post /api/chat/user_top (userTopRequest) returns (userTopResponse) // 好友置顶
}

// goctl api go --api chat.api --dir .

```

#### 聊天记录服务

```go
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
	Msg       common.Message    `json:"msg"`
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

```

#### 最近会话记录

```go
package logic

import (
	"context"
	"fmt"
	"server/chat/api/internal/svc"
	"server/chat/api/internal/types"
	"server/common/list_query"
	"server/common/models"
	model "server/models/chat"
	user_rpc "server/user/rpc/userservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatSessionLogic {
	return &ChatSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Data struct {
	SU         uint   `gorm:"column:sU"`
	RU         uint   `gorm:"column:rU"`
	MaxDate    string `gorm:"column:maxDate"`
	MaxPreview string `gorm:"column:maxPreview"`
	IsTop      bool   `gorm:"column:isTop"`
}

func (l *ChatSessionLogic) ChatSession(req *types.ChatSessionRequest) (resp *types.ChatSessionResponse, err error) {
	isTop := fmt.Sprintf(" if((select 1 from top_user_models where user_id = %d and (top_user_id = sU or top_user_id = rU) limit 1), 1, 0)  as isTop", req.UserID)
	var friendIDList []uint
	// 找出所有好友信息
	r, err := l.svcCtx.UserRpc.FriendList(context.Background(), &user_rpc.FriendListRequest{
		UserId: uint32(req.UserID),
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	friendInfoMap := r.UserInfoMap
	// 将查出的会话信息映射到Data结构体
	// 先获取到 好友ID列表
	friendIDList = make([]uint, 0)
	for id, _ := range friendInfoMap {
		friendIDList = append(friendIDList, uint(id))
	}
	topIDs := make([]uint, 0)
	// 在置顶表里查询置顶的好友ID
	err = l.svcCtx.DB.Model(&model.TopUserModel{}).Where("user_id = ?", req.UserID).Pluck("top_user_id", &topIDs).Error
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	chats, count, err := list_query.ListQuery(l.svcCtx.DB, Data{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.PageSize,
		},
		Sort: "isTop desc, maxDate desc",
		Table: func() (string, any) {
			return "(?) as u", l.svcCtx.DB.Model(&model.ChatModel{}).
				Select("least(send_user_id, rev_user_id) as sU ,greatest(send_user_id, rev_user_id) as rU, max(created_at) as maxDate",
					fmt.Sprintf("(select msg_preview from chat_models  where ((send_user_id = sU and rev_user_id = rU) or (send_user_id = rU and rev_user_id = sU)) and id not in (select chat_id from user_chat_delete_models where user_id = %d) order by created_at desc  limit 1) as maxPreview", req.UserID),
					isTop).
				Where("(send_user_id = ? or rev_user_id = ?) and id not in (select chat_id from user_chat_delete_models where user_id = ?) and (send_user_id = ? and rev_user_id in ?) or (rev_user_id = ? and send_user_id in ?)",
					req.UserID, req.UserID, req.UserID, req.UserID, friendIDList, req.UserID, friendIDList).
				Group("least(send_user_id, rev_user_id)").
				Group("greatest(send_user_id, rev_user_id)")
		},
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	chatSessions := make([]types.ChatSession, 0)
	// 处理会话, 将会话的用户信息附上
	for _, c := range chats {
		friendID := c.SU
		if c.RU != req.UserID {
			friendID = c.RU
		}
		session := types.ChatSession{
			UserID:     friendID, // 好友ID
			Avatar:     friendInfoMap[uint32(friendID)].Avatar,
			Nickname:   friendInfoMap[uint32(friendID)].NickName,
			CreatedAt:  c.MaxDate,
			MsgPreview: c.MaxPreview,
			IsTop:      c.IsTop,
			IsOnline:   friendInfoMap[uint32(friendID)].IsOnline,
		}
		chatSessions = append(chatSessions, session)
	}
	return &types.ChatSessionResponse{
		List:  chatSessions,
		Count: int(count),
	}, nil
}

```

#### 用户聊天功能（重要）

##### 视频聊天

> 视频聊天使用WebRTC 进行实时视频聊天，后端服务器主要用来建立 WebRTC 连接，因为很多浏览器已经内置了 WebRTC技术

**前置内容**

```go
type VideoCallMessage struct {
	StartTime  time.Time `json:"startTime"`  // 通话开始时间
	EndTime    time.Time `json:"endTime"`    // 通话结束时间
	EndReason  int8      `json:"endReason"`  // 通话结束原因 
	ClientFlag int8      `json:"clientFlag"` // 标识客户端弹框的模式
	ServerFlag int8      `json:"serverFlag"` // 服务器处理逻辑
	Msg        string    `json:"msg"`        // 提示消息
	Type       string    `json:"type"`       // WebRTC 消息类型 create_offer create_answer create_candidate 等
	Data       any       `json:"data"`       // 附加数据如 offer answer candidate
}
```

 假设 A -> B 发起通话请求, Flag情况分析如下

<!-- tabs:start --> 

#### **ClientFlag**

> ClientFlag（客户端接收） 标识客户端弹框的模式

1 ： 客户端展示一个等待对方接听的弹框（自己发起的通话请求）

2 ： 客户端展示等待接听的弹框（对方发来的通话请求）

3 : 客户端展示对方未接听的弹框 

4 ：客户端展示对方已拒绝的弹框

5 :  客户端展示自己挂断的弹框

6 ：客户端准备发送 WebRTC 的 offer

7 :  客户端展示对方已挂断的弹框

#### **ServerFlag**

> ServerFlag（服务端接收） 判断消息模式

0 ： 用户发起通话请求  					     --> **处理**： 服务器向 A 发送 flag：1 ，向 B 发送 flag：2 

1 ： 通话发起者挂断，通话没有建立就挂断了  		--> **处理:**  [服务器向 A 发送 flag：3 对方未接通] [向 B 发送flag：4 未接听 ] 通话结束，消息入库

2 :  对方拒绝通话请求, 挂断弹框由对方前端处理 	   --> **处理:** 服务器向 A 发送 flag：4，通话结束，消息入库

3 ： 对方接受通话请求						 --> **处理:** 服务器向 A 发送 flag：5，A 收到flag：5 需准备发送 offer



<!-- tabs:end --> 

**主要步骤**

 假设 A -> B 发起通话请求, 有如下情况

1. A 发起请求, 等待B接听, B还未接听, A主动挂电话
2. A 发起请求, B主动挂电话
3. A 发起请求, B接受了通话
4. A - B 通话中, A主动挂了电话
5. A - B 通话中, B主动挂了电话



