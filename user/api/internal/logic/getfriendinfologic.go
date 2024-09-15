package logic

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	user "server/models/user"
	"server/user/api/internal/svc"
	"server/user/api/internal/types"
	user_rpc "server/user/rpc/userservice"
)

type GetFriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInfoLogic {
	return &GetFriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendInfoLogic) GetFriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	// 先判断两者是否是好友
	friendship := user.FriendshipModel{}
	isFriend := friendship.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID)
	if !isFriend {
		return nil, errors.New("你们还不是好友呢！")
	}
	// 如果是好友则获取好友信息， friendship 已经获取到了
	friendInfo := user.UserModel{}
	info, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: uint32(req.FriendID),
	})
	// info 里面的 password 已经去除
	json.Unmarshal(info.Data, &friendInfo)
	return &types.FriendInfoResponse{
		NickName: friendInfo.Nickname,
		Avatar:   friendInfo.Avatar,
		Alias:    friendship.Alias,
		Abstract: friendInfo.Abstract,
	}, nil
}
