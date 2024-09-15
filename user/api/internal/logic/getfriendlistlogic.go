package logic

import (
	"context"
	"server/common/list_query"
	common "server/common/models"
	models "server/models/user"

	"server/user/api/internal/svc"
	"server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.GetFriendListRequest) (resp *types.GetFriendListResponse, err error) {
	// 先获取 friendShip 表中的数据

	friendShipList, count, err := list_query.ListQuery(l.svcCtx.DB, models.FriendshipModel{}, list_query.Option{
		PageInfo: common.PageInfo{
			Page:  int(req.Page),
			Limit: int(req.PageSize),
		},
		Where: l.svcCtx.DB.Where("user_id = ? or friend_id = ?", req.UserID, req.UserID), // 找 UserID 的好友列表
	})
	if err != nil {
		return nil, err
	}
	if count <= 0 {
		return nil, nil
	}
	idList := []uint{}
	for _, v := range friendShipList {
		if v.UserID == req.UserID {
			idList = append(idList, v.FriendID)
		} else {
			idList = append(idList, v.UserID)
		}
	}
	// 再查user
	friendList, count, err := list_query.ListQuery(l.svcCtx.DB, models.UserModel{}, list_query.Option{
		PageInfo: common.PageInfo{
			Page:  -1,
			Limit: -1,
		},
		Where: l.svcCtx.DB.Where("id in ?", idList),
	})
	list := make([]types.FriendInfoResponse, count)
	for i, v := range friendList {
		list[i] = types.FriendInfoResponse{
			NickName: v.Nickname,
			Avatar:   v.Avatar,
			Abstract: v.Abstract,
			Alias:    friendShipList[i].Alias,
		}
	}
	return &types.GetFriendListResponse{
		List:  list,
		Count: uint(count),
	}, nil
}
