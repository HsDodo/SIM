package logic

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"server/auth/api/internal/svc"
	"server/auth/api/internal/types"
	"server/common/logger"
	"server/utils/jwt"
)

type AuthenticationLogic struct {
	Logger *logrus.Entry
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {

	return &AuthenticationLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp types.AuthenticationReponse, err error) {
	//token不能为空
	if token == "" {
		return types.AuthenticationReponse{}, errors.New("token不能为空")
	}
	//解析token
	payload, err := jwt.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return types.AuthenticationReponse{}, err
	}
	//判断token是否在redis黑名单中
	_, err = l.svcCtx.Redis.Get(l.ctx, token).Result()
	// 如果token在redis黑名单中，返回错误
	if err == nil { //查到了，说明token已注销，判断是不是同一个token
		return types.AuthenticationReponse{}, errors.New("token已注销")
	}
	return types.AuthenticationReponse{
		UserId:   payload.UserID,
		NickName: payload.NickName,
		Role:     payload.Role,
	}, nil
}
