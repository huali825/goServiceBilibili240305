package ioc

import (
	"go20240218/01webook/internal/service/sms"
	"go20240218/01webook/internal/service/sms/memory"
)

func InitSMSService() sms.Service {
	return memory.NewService()
	// 如果有需要，就可以用这个
	//return initTencentSMSService()
}
