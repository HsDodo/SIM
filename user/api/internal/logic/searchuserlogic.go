package logic

import (
	"context"
	"server/common/list_query"
	common "server/common/models"
	models "server/models/user"
	"server/user/api/internal/svc"
	"server/user/api/internal/types"
	"server/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUserLogic) SearchUser(req *types.SearchUserRequest) (resp *types.SearchUserResponse, err error) {

	// 首先查找那些在线的用户
	userList, count, err := list_query.ListQuery(l.svcCtx.DB, models.UserModel{}, list_query.Option{
		PageInfo: common.PageInfo{
			Page:  req.Page,
			Limit: req.PageSize,
		},
		Preload: []string{"UserConf"},
		Joins:   "left join user_conf_models ucm on ucm.user_id = user_models.id",
		Where:   l.svcCtx.DB.Where("(ucm.search_user <> 0 or ucm.search_user is not null) and (ucm.search_user = 1 and user_models.id = ?) or (ucm.search_user = 2 and (user_models.id = ? or user_models.nickname like ?))", req.Key, req.Key, "%"+req.Key+"%"),
	})
	if err != nil {
		return nil, err
	}

	// 查出来了用户列表, 进行封装
	var searchInfoList []types.SearchInfo
	friendIDList, err := models.GetFriendshipIDs(l.svcCtx.DB, req.UserID)
	for _, user := range userList {
		isFriend := utils.InIDsList(friendIDList, user.ID)
		searchInfoList = append(searchInfoList, types.SearchInfo{
			UserID:   user.ID,
			NickName: user.Nickname,
			Avatar:   user.Avatar,
			Abstract: user.Abstract,
			IsFriend: isFriend,
		})
	}
	return &types.SearchUserResponse{
		List:  searchInfoList,
		Count: uint(count),
	}, nil
}
