package logic

import (
	"context"
	"server/common/list_query"
	"server/common/models"
	group_models "server/models/group"

	"server/group/api/internal/svc"
	"server/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupMyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupMyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupMyLogic {
	return &GroupMyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupMyLogic) GroupMy(req *types.GroupMyRequest) (resp *types.GroupMyListResponse, err error) {
	// Note: 我创建的群聊列表 或者加入的群聊列表
	// 查群id列表
	var groupIDList []uint
	query := l.svcCtx.DB.Model(&group_models.GroupMemberModel{}).Where("user_id = ?", req.UserID)
	if req.Mode == 1 {
		// 我创建的群聊
		query.Where("role = ?", 2)
	}
	query.Pluck("group_id", &groupIDList)

	groupList, count, _ := list_query.ListQuery(l.svcCtx.DB, group_models.GroupModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.PageSize,
		},
		Preload: []string{"MemberList"},
		Where:   l.svcCtx.DB.Where("id in ?", groupIDList),
	})
	resp = new(types.GroupMyListResponse)
	for _, group := range groupList {

		var role int8
		for _, memberModel := range group.MemberList {
			if memberModel.UserID == req.UserID {
				role = memberModel.Role
			}
		}

		resp.List = append(resp.List, types.GroupMyResponse{
			GroupID:          group.ID,
			GroupName:        group.GroupName,
			GroupAvatar:      group.GroupAvatar,
			GroupMemberCount: len(group.MemberList),
			Role:             role,
			Mode:             req.Mode,
		})
	}

	resp.Count = int(count)
	return
}
