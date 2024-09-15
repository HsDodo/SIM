package logic

import (
	"context"
	"errors"
	models "server/models/user"

	"server/user/api/internal/svc"
	"server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFriendAliasLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFriendAliasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFriendAliasLogic {
	return &UpdateFriendAliasLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateFriendAliasLogic) UpdateFriendAlias(req *types.UpdateFriendAliasRequest) (resp *types.UpdateFriendAliasResponse, err error) {
	// 先判断两者是否是好友
	friendShip := models.FriendshipModel{}
	isFriend := friendShip.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID)
	if !isFriend {
		return nil, errors.New("不是好友关系")
	}
	err = l.svcCtx.DB.Model(&friendShip).Update("alias", req.Alias).Error
	return &types.UpdateFriendAliasResponse{}, err
}
