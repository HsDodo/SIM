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
