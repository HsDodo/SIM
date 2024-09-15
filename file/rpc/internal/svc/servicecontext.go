package svc

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"server/common/logger"
	"server/core"
	"server/file/rpc/internal/config"
	user_rpc "server/user/rpc/userservice"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := c.Mysql.DataSource
	db := core.InitGorm(dsn)
	redisClient, err := core.InitRedis(c.RedisConf.Addr, c.RedisConf.Pwd, c.RedisConf.Db)
	if err != nil {
		logger.Fatalf("redis初始化失败!: %v", err)
	}
	return &ServiceContext{
		Config: c,
		DB:     db,
		Redis:  redisClient,
	}
}
