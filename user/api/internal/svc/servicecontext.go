package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	chat_rpc "server/chat/rpc/chatservice"
	"server/common/logger"
	"server/common/zrpc_interceptor"
	"server/core"
	"server/user/api/internal/config"
	user_rpc "server/user/rpc/userservice"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UserService
	ChatRpc chat_rpc.ChatService
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
		ChatRpc: chat_rpc.NewChatService(zrpc.MustNewClient(c.ChatRpc, zrpc.WithUnaryClientInterceptor(zrpc_interceptor.ClientInfoInterceptor))),
	}
}
