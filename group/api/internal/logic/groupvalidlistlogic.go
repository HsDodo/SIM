package logic

import (
	"context"
	"server/common/list_query"
	"server/common/models"
	group_models "server/models/group"
	user_rpc "server/user/rpc/userservice"

	"server/group/api/internal/svc"
	"server/group/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupValidListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupValidListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupValidListLogic {
	return &GroupValidListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupValidListLogic) GroupValidList(req *types.GroupValidListRequest) (resp *types.GroupValidListResponse, err error) {
	// 查找我管理的群的id列表
	var groupIDList []uint
	err = l.svcCtx.DB.Model(group_models.GroupMemberModel{}).Where("user_id = ? and (role = 1 or role = 2)", req.UserID).Select("group_id").Scan(&groupIDList).Error
	if err != nil {
		logx.Error(err)
		return
	}
	var groupMap = map[uint]bool{}
	for _, id := range groupIDList {
		groupMap[id] = true
	}

	verifyList, count, err := list_query.ListQuery(l.svcCtx.DB, group_models.GroupVerifyModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.PageSize,
		},
		Preload: []string{"GroupModel"},
		Where:   l.svcCtx.DB.Where("group_id in ? or user_id = ?", groupIDList, req.UserID),
	})

	var userIDList []uint32
	for _, verify := range verifyList {
		userIDList = append(userIDList, uint32(verify.UserID))
	}
	// 查这些人的信息
	userInfoMapData, err := l.svcCtx.UserRpc.UserListInfo(l.ctx, &user_rpc.UserListInfoRequest{
		UserIDs: userIDList,
	})

	var userInfoMap = map[uint32]*user_rpc.UserInfo{}
	userInfoMap = userInfoMapData.UserInfoMap

	if err != nil {
		logx.Error(err)
		return
	}

	for _, verify := range verifyList {
		verifyInfo := types.GroupValidInfoResponse{
			ID:                 verify.ID,
			GroupID:            verify.GroupID,
			UserID:             verify.UserID,
			AdditionalMessages: verify.AdditionalMessages,
			Status:             int8(verify.Status),
			Type:               verify.Type,
			CreatedAt:          verify.CreatedAt.String(),
			GroupName:          verify.GroupModel.GroupName,
			Avatar:             verify.GroupModel.GroupAvatar,
			Flag:               "send",
			UserAvatar:         userInfoMap[uint32(verify.UserID)].Avatar,
			UserNickname:       userInfoMap[uint32(verify.UserID)].NickName,
		}
		if groupMap[verify.GroupID] {
			verifyInfo.Flag = "rev"
		}
		if verify.VerificationQuestion != nil {
			verifyInfo.VerificationQuestion = &types.VerificationQuestion{
				Problem1: verify.VerificationQuestion.Problem1,
				Problem2: verify.VerificationQuestion.Problem2,
				Problem3: verify.VerificationQuestion.Problem3,
				Answer1:  verify.VerificationQuestion.Answer1,
				Answer2:  verify.VerificationQuestion.Answer2,
				Answer3:  verify.VerificationQuestion.Answer3,
			}
		}
		resp.List = append(resp.List, verifyInfo)
	}
	resp.Count = int(count)
	return
}
