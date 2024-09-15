package logic

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"server/common/logger"

	"server/auth/api/internal/svc"
	"server/auth/api/internal/types"
	"server/models/user"
	jwt "server/utils/jwt"
)

type LoginLogic struct {
	Logger *logrus.Entry
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	//查询用户
	user := models.UserModel{}
	err = l.svcCtx.DB.Model(&models.UserModel{}).Where("nickname = ?", req.UserName).Take(&user).Error
	if err != nil {
		err = errors.New("用户名或密码错误")
		logger.LogError(err)
		return
	}
	//判断密码是否正确
	err = bcrypt.CompareHashAndPassword([]byte(user.Pwd), []byte(req.Password))
	if err != nil {
		e := errors.New("密码错误")
		logger.LogError(err)
		return nil, e
	}
	//登录成功，进行处理
	token, err := jwt.GenJwtToken(jwt.PayLoad{
		UserID:   user.ID,
		NickName: user.Nickname,
		Role:     user.Role,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		err = errors.New("生成token失败！")
		return
	}
	logger.LogWithFields(logrus.Fields{
		"userName": user.Nickname,
		"userId":   user.ID,
		"pwd":      req.Password,
	}).Info("登录成功！")

	//将user存入svcCtx中
	l.svcCtx.User = user

	return &types.LoginResponse{
		Token: token,
	}, nil
}
