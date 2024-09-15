package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"server/chat/api/internal/config"
	"server/common/logger"
	"server/common/zrpc_interceptor"
	"server/core"
	file_rpc "server/file/rpc/fileservice"
	user_rpc "server/user/rpc/userservice"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UserService
	FileRpc file_rpc.FileService
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.Mysql.DataSource
	db := core.InitGorm(dsn)
	redisClient, err := core.InitRedis(c.Redis.Addr, c.Redis.Pwd, c.Redis.DB)
	if err != nil {
		logger.Fatalf("redis初始化失败!: %v", err)
	}
	return &ServiceContext{
		Config:  c,
		DB:      db,
		Redis:   redisClient,
		UserRpc: user_rpc.NewUserService(zrpc.MustNewClient(c.UserRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
		FileRpc: file_rpc.NewFileService(zrpc.MustNewClient(c.FileRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
	}
}
