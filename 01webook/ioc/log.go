package ioc

import (
	"go.uber.org/zap"
	"go20240218/01webook/pkg/logger"
)

func InitLogger() logger.LoggerV1 {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
