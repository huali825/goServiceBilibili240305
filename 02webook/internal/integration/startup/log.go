package startup

import (
	"go20240218/02webook/pkg/logger"
)

func InitLog() logger.LoggerV1 {
	return &logger.NopLogger{}
}
