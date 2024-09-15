package logic

import (
	"context"
	"server/user/rpc/internal/svc"
	"server/user/rpc/proto"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOnlineListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOnlineListLogic {
	return &UserOnlineListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOnlineListLogic) UserOnlineList(in *proto.UserOnlineListRequest) (*proto.UserOnlineListResponse, error) {
	// Note: 返回在 in 中的用户列表中在线的用户
	var onlineUserIDs []uint32
	onlineUserMap := l.svcCtx.Redis.HGetAll(context.Background(), "online").Val()

	for _, userID := range in.UserIds {
		//判断在redis中是否登录了
		if _, ok := onlineUserMap[strconv.Itoa(int(userID))]; ok {
			onlineUserIDs = append(onlineUserIDs, userID)
		}
	}
	return &proto.UserOnlineListResponse{
		UserIds: onlineUserIDs,
	}, nil
}
