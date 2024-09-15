package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// 微服务配置
	UserRpc zrpc.RpcClientConf
	FileRpc zrpc.RpcClientConf
	ChatRpc zrpc.RpcClientConf
	Etcd    struct {
		Hosts []string
	}
	Mysql struct {
		DataSource string
	}
	Redis struct {
		Addr string
		Pwd  string
		DB   int
	}
}
