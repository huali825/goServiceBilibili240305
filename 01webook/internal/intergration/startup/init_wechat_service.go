package startup

import (
	"go20240218/01webook/internal/service/oauth2/wechat"
	"go20240218/01webook/pkg/logger"
)

// InitPhantomWechatService 没啥用的虚拟的 wechatService
func InitPhantomWechatService(l logger.LoggerV1) wechat.Service {
	return wechat.NewService("", "", l)
}
