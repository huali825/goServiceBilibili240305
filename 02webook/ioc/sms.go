package ioc

import (
	"github.com/redis/go-redis/v9"
	"go20240218/02webook/internal/service/sms"
	"go20240218/02webook/internal/service/sms/memory"
)

func InitSMSService(cmd redis.Cmdable) sms.Service {
	// 换内存，还是换别的
	//svc := ratelimit.NewRatelimitSMSService(memory.NewService(),
	//	limiter.NewRedisSlidingWindowLimiter(cmd, time.Second, 100))
	//return retryable.NewService(svc, 3)
	return memory.NewService()
}
