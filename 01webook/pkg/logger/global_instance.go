package logger

import "sync"

var gl LoggerV1
var lMutex sync.RWMutex

func SetGlobalLogger(l LoggerV1) {
	lMutex.Lock()
	defer lMutex.Unlock()
	gl = l
}

func L() LoggerV1 {
	lMutex.RLock()
	g := gl
	lMutex.RUnlock()
	return g
}

// GL 如果对线程安全 无感就这么搞 但是不这么搞 一般是测试的时候用
var GL LoggerV1 = &NopLogger{}
