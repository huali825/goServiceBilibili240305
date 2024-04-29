package ioc

import (
	"go20240218/01webook/internal/service/sms"
	"go20240218/01webook/internal/service/sms/memory"
)

// InitSMSService 短信服务随时替换
func InitSMSService() sms.Service {
	return memory.NewService()
	// 如果有需要，就可以用这个
	//return initTencentSMSService()
}
