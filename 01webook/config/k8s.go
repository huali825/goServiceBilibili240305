//go:build k8s

// 使用 k8s 这个编译标签
package config

var Config = config{
	DB: DBConfig{
		// 本地连接
		DSN: "root:root@tcp(webook-mysql:11309)/webook",
	},
	Redis: RedisConfig{
		Addr: "webook-live-redis:11479",
	},
	HS: HelloString{
		str: "k8s hello",
	},
}
