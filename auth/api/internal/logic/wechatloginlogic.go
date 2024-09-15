package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"server/auth/api/internal/svc"
)

type WechatLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatLoginLogic {
	return &WechatLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatLoginLogic) WechatLogin() (resp bool, err error) {
	// todo: add your logic here and delete this line

	return
}
