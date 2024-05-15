package ioc

import (
	"github.com/redis/go-redis/v9"
	"go20240218/01webook/internal/service/sms"
	"go20240218/01webook/internal/service/sms/memory"
	"go20240218/01webook/internal/service/sms/ratelimit"
	"go20240218/01webook/internal/service/sms/retryable"
	ratelimit1 "go20240218/01webook/pkg/ratelimit"
	"time"
)

// InitSMSService 短信服务随时替换
func InitSMSService(cmd redis.Cmdable) sms.Service {
	//return memory.NewService()
	// 如果有需要，就可以用这个
	//return initTencentSMSService()

	svc := ratelimit.NewRatelimitSMSService(memory.NewService(),
		ratelimit1.NewRedisSlidingWindowLimiter(cmd, time.Second, 100))
	return retryable.NewService(svc, 3)
}
