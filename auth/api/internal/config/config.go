package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// 微服务配置
	UserRpc zrpc.RpcClientConf
	Etcd    struct {
		Hosts []string
	}
	Mysql struct {
		DataSource string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	SystemEnv struct {
		LogLevel string
	}
	Redis struct {
		Addr string
		Pwd  string
		DB   int
	}
	OpenLoginList OpenLoginList
}

type Endpoint struct {
	TokenURL    string
	AuthURL     string
	UserInfoURL string
}

type OpenLoginList struct {
	Wechat OpenLogin
}

type OpenLogin struct {
	Appid       string
	AppSecret   string
	Icon        string
	RedirectURI string
	Endpoint    Endpoint
	Scopes      []string
}
