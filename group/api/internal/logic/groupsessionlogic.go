package logic

import (
	"context"
	"fmt"
	"server/common/list_query"
	"server/common/models"
	group_models "server/models/group"

	"server/group/api/internal/svc"
	"server/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupSessionLogic {
	return &GroupSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type SessionData struct {
	GroupID       uint   `gorm:"column:groupID"`
	NewMsgDate    string `gorm:"column:newMsgDate"`
	NewMsgPreview string `gorm:"column:newMsgPreview"`
	IsTop         bool   `gorm:"column:isTop"`
}

func (l *GroupSessionLogic) GroupSession(req *types.GroupSessionRequest) (resp *types.GroupSessionListResponse, err error) {
	// 先查我有哪些群
	var userGroupIDList []uint
	err = l.svcCtx.DB.Model(group_models.GroupMemberModel{}).
		Where("user_id = ?", req.UserID).Pluck("group_id", &userGroupIDList).Error
	if err != nil {
		return nil, err
	}
	column := fmt.Sprintf(" (if((select 1 from group_user_top_models where user_id = %d and group_user_top_models.group_id = group_msg_models.group_id), 1, 0)) as isTop", req.UserID)

	// 查哪些聊天记录是被删掉的
	var msgDeleteIDList []uint
	l.svcCtx.DB.Model(group_models.GroupUserMsgDeleteModel{}).Where("group_id in ?", userGroupIDList).Pluck("msg_id", &msgDeleteIDList)
	query := l.svcCtx.DB.Where("group_id in (?)", userGroupIDList)
	if len(msgDeleteIDList) > 0 {
		query.Where("id not in ?", msgDeleteIDList)
	}

	sessionList, count, _ := list_query.ListQuery(l.svcCtx.DB, SessionData{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.PageSize,
		},
		Sort:  "isTop desc, newMsgDate desc",
		Debug: true,
		Table: func() (string, any) {
			return "(?) as u", l.svcCtx.DB.Model(&group_models.GroupMsgModel{}).
				Select("group_id as groupID",
					"max(created_at) as newMsgDate",
					column,
					"(select msg_preview from group_msg_models as g where g.group_id = groupID order by g.created_at desc limit 1)  as newMsgPreview").
				Where(query).
				Group("group_id")
		},
	})

	var groupIDList []uint
	for _, session := range sessionList {
		groupIDList = append(groupIDList, session.GroupID)
	}

	var groupList []group_models.GroupModel
	err = l.svcCtx.DB.Find(&groupList, "id in ?", groupIDList).Error
	if err != nil {
		return nil, err
	}
	groupMap := make(map[uint]group_models.GroupModel)
	for _, group := range groupList {
		groupMap[group.ID] = group
	}

	resp = new(types.GroupSessionListResponse)
	for _, session := range sessionList {
		resp.List = append(resp.List, types.GroupSessionResponse{
			GroupID:       session.GroupID,
			GroupName:     groupMap[session.GroupID].GroupName,
			GroupAvatar:   groupMap[session.GroupID].GroupAvatar,
			NewMsgDate:    session.NewMsgDate,
			NewMsgPreview: session.NewMsgPreview,
			IsTop:         session.IsTop,
		})
	}
	resp.Count = int(count)

	return
}
