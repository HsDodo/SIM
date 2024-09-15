package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	//Mysql 数据库配置
	Mysql struct {
		DataSource string
	}
	//Redis 缓存配置
	RedisConf struct {
		Addr string
		Pwd  string
		Db   int
	}
}
