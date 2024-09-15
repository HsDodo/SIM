package logic

import (
	"context"
	"encoding/json"
	models "server/models/chat"

	"server/chat/rpc/internal/svc"
	"server/chat/rpc/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserChatLogic {
	return &UserChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserChatLogic) UserChat(in *proto.UserChatRequest) (*proto.NoneResponse, error) {
	// todo: 创建聊天记录
	var chat models.ChatModel
	json.Unmarshal(in.ChatMsg, &chat)
	l.svcCtx.DB.Create(&chat)

	return &proto.NoneResponse{}, nil
}
