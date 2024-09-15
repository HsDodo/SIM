package core

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitRedis(addr string, pwd string, db int) (*redis.Client, error) {
	// 初始化redis
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd, // no password set
		DB:       db,  // use default DB
	})
	// 测试连接
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	// 返回client
	return client, nil
}
