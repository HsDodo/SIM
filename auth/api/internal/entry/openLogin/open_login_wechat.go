package openLogin

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type WechatInfo struct {
	Nickname    string   `json:"nickname"`   // 昵称
	Sex         int      `json:"sex"`        // 性别
	HeadImgUrl  string   `json:"headimgurl"` // 头像大图
	OpenID      string   `json:"openid"`
	Privilege   []string `json:"privilege"`   // 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	UnionID     string   `json:"unionid"`     // 用户统一标识
	Province    string   `json:"province"`    // 省份
	City        string   `json:"city"`        // 城市
	Openid_list string   `json:"openid_list"` // openid列表
	Address     string   `json:"address"`     // 地址
	Email       string   `json:"email"`       // 邮箱
	Country     string   `json:"country"`     // 国家
}

type WechatLogin struct {
	WechatConfig
	code        string
	accessToken string
	openID      string
}

type WechatConfig struct {
	AppID     string
	AppSecret string
	TokenURL  string
	InfoURL   string
	Redirect  string
}

func NewWechatLogin(code string, conf WechatConfig) (wechatInfo WechatInfo, err error) {
	wechatLogin := &WechatLogin{
		WechatConfig: conf,
		code:         code,
	}
	err = wechatLogin.GetAccessTokenAndOpenId() //获取access_token和open_id
	if err != nil {
		return wechatInfo, err
	}
	wechatInfo, err = wechatLogin.GetUserInfo() //获取用户信息
	if err != nil {
		return wechatInfo, err
	}
	wechatInfo.OpenID = wechatLogin.openID
	return wechatInfo, nil
}

func (w *WechatLogin) GetAccessTokenAndOpenId() (err error) {
	params := url.Values{}
	params.Add("grant_type", "authorization_code") // 固定值 authorization_code 表示使用授权码模式
	params.Add("appid", w.AppID)
	params.Add("secret", w.AppSecret)
	params.Add("code", w.code)
	// 获取access_token ,不需要携带回调地址, 官方样例: https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&lcode=CODE&grant _type=authorization_code
	u, err := url.Parse(w.TokenURL)
	if err != nil {
		return err
	}
	u.RawQuery = params.Encode()
	resp, err := http.Get(u.String()) //获取token
	if err != nil {
		return err
	}
	m := make(map[string]interface{})
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	w.accessToken = m["access_token"].(string) //将access_token赋值给结构体
	w.openID = m["openid"].(string)            //将open_id赋值给结构体
	return nil
}

func (w *WechatLogin) GetUserInfo() (WechatInfo, error) {
	// http：GET（请使用https协议）：
	// 获取用户信息, 官方样例: https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
	// 需要参数: access_token, openid, lang(默认为zh_CN),
	userInfo := WechatInfo{}
	params := url.Values{}
	params.Add("access_token", w.accessToken)
	params.Add("openid", w.openID)
	params.Add("lang", "zh_CN")
	params.Add("scope", "snsapi_userinfo,snsapi_privateinfo,snsapi_friend")

	u, err := url.Parse(w.InfoURL) //userInfo请求地址
	if err != nil {
		return WechatInfo{}, err
	}
	u.RawQuery = params.Encode()
	resp, err := http.Get(u.String()) //获取用户信息 , 返回json数据包
	if err != nil {
		return WechatInfo{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&userInfo) //解析json数据
	if err != nil {
		return WechatInfo{}, err
	}
	return userInfo, nil
}
