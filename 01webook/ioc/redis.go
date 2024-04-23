package ioc

import (
	"github.com/redis/go-redis/v9"
	"go20240218/01webook/config"
)

func InitRedis() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
}
