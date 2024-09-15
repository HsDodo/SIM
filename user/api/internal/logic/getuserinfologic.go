package logic

import (
	"context"
	user_rpc "server/user/rpc/userservice"

	"server/user/api/internal/svc"
	"server/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.UserInfoResponse, err error) {
	// 获取用户信息逻辑
	info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user_rpc.UserInfoRequest{
		UserId: uint32(req.UserID),
	})
	if err != nil {
		l.Errorf("RPC 获取用户信息失败: %v", err)
		return nil, err
	}

	return &types.UserInfoResponse{Data: info.Data}, nil
}
