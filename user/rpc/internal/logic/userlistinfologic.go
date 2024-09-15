package logic

import (
	"context"
	"errors"
	models "server/models/user"

	"server/user/rpc/internal/svc"
	"server/user/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListInfoLogic {
	return &UserListInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListInfoLogic) UserListInfo(req *proto.UserListInfoRequest) (*proto.UserListInfoResponse, error) {
	// 按照用户ID列表查询用户信息
	userList := make([]models.UserModel, 0)
	if len(req.UserIDs) > 0 {
		err := l.svcCtx.DB.Preload("UserConf").Where("id in (?)", req.UserIDs).Find(&userList).Error
		if err != nil {
			return nil, errors.New("查询用户信息失败")
		}
	}
	userInfoMap := make(map[uint32]*proto.UserInfo, 0)
	for _, user := range userList {
		userInfoMap[uint32(user.ID)] = &proto.UserInfo{
			UserId:   uint32(user.ID),
			NickName: user.Nickname,
			Avatar:   user.Avatar,
			IsOnline: user.UserConf.OnlineStatus,
		}
	}
	return &proto.UserListInfoResponse{
		UserInfoMap: userInfoMap,
	}, nil
}
