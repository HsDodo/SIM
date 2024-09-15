package logic

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"server/auth/api/internal/svc"
	"server/common/logger"
	"server/models/user"
	"server/utils/jwt"
	"time"
)

type LogoutLogic struct {
	Logger *logrus.Entry
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp string, err error) {
	//解析token
	//token解析成功后，将token存入redis黑名单
	//将当前用户信息从ServiceContext中删除
	//返回注销成功
	claims, err := jwt.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return "", err
	}
	exp := claims.ExpiresAt.Time.Sub(time.Now())
	//将token存入redis黑名单，在token过期时间后自动删除
	_, err = l.svcCtx.Redis.SetNX(l.ctx, token, 1, exp).Result()
	if err != nil {
		return "", errors.New("注销失败: redis 存入token失败！")
	}

	logger.LogWithFields(logrus.Fields{
		"userId":   claims.ID,
		"userName": claims.NickName,
	}).Println("注销成功！")
	username := l.svcCtx.User.Nickname
	l.svcCtx.User = models.UserModel{}
	return "注销成功: " + username, nil
}
