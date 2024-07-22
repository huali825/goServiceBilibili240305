package loggerxx

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger(l *zap.Logger) {
	Logger = l
}

// InitLoggerV1 main 函数调用一下
func InitLoggerV1() {
	Logger, _ = zap.NewDevelopment()
}

//var SecureLogger *zap.Logger
