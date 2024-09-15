package logic

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/auth/api/internal/entry/openLogin"
	"server/auth/api/internal/svc"
	"server/auth/api/internal/types"
	"server/common/logger"
	users "server/models/user"
	user_proto "server/user/rpc/proto"
	"server/utils/jwt"
)

type OpenLoginLogic struct {
	Logger *logrus.Entry
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpenLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenLoginLogic {
	return &OpenLoginLogic{
		Logger: logger.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *OpenLoginLogic) OpenLogin() (resp *types.LoginResponse, err error) {
	// todo: 这里是用来处理回调逻辑的, 第三方登录服务器携带code回调到这里
	loginType := l.svcCtx.Redis.Get(l.ctx, "OpenloginType").Val()
	var user users.UserModel
	loginType = "wechat"
	switch loginType {
	case "wechat":
		// 1. 获取code
		r := l.ctx.Value("Request").(*http.Request)
		code := r.URL.Query().Get("code") //code随着回调地址返回了
		wechatConf := l.svcCtx.Config.OpenLoginList.Wechat
		// 2. 通过code获取用户信息, NewWechatLogin 封装了获取用户信息的逻辑
		userInfo, err := openLogin.NewWechatLogin(code, openLogin.WechatConfig{
			AppID:     wechatConf.Appid,
			AppSecret: wechatConf.AppSecret,
			Redirect:  wechatConf.RedirectURI,
			TokenURL:  wechatConf.Endpoint.TokenURL,
			InfoURL:   wechatConf.Endpoint.UserInfoURL,
		})
		if err != nil {
			return nil, err
		}

		// 3. 获取了用户信息之后, 可以进行业务逻辑处理, 比如判断用户是否存在, 不存在则创建用户
		user = users.UserModel{}
		err = l.svcCtx.DB.Model(&users.UserModel{}).Where("open_id = ?", userInfo.OpenID).First(&user).Error
		if err != nil {
			// 调用 rpc 创建用户
			_, err := l.svcCtx.UserRpc.CreateUser(l.ctx, &user_proto.UserCreateRequest{
				OpenId:       userInfo.OpenID,
				NickName:     userInfo.Nickname,
				Avatar:       userInfo.HeadImgUrl,
				RegisterType: "wechat",
			})

			if err != nil {
				logger.LogErrorStr("创建用户失败")
				return nil, err
			} else {
				logger.LogWithFields(logrus.Fields{
					"OpenId":       userInfo.OpenID,
					"Nickname":     user.Nickname,
					"Avatar":       userInfo.HeadImgUrl,
					"RegisterType": "wechat",
				}).Println("注册成功！")
			}
		}
	default:
	}
	// 登录之后颁发token
	//登录成功，进行处理
	token, err := jwt.GenJwtToken(jwt.PayLoad{
		UserID:   user.ID,
		NickName: user.Nickname,
	}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		err = errors.New("生成token失败！")
		return
	}
	logger.LogWithFields(logrus.Fields{
		"userId":    user.ID,
		"userName":  user.Nickname,
		"LoginType": loginType,
	}).Println("登录成功！")
	// 将用户存入svcCtx中
	l.svcCtx.User = user
	return &types.LoginResponse{
		Token: token,
	}, nil
}
