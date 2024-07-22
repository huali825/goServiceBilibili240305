//go:build !k8s

// asdsf go:build dev
// sdd go:build test
// dsf 34

// 没有k8s 这个编译标签
package config

var Config = config{
	DB: DBConfig{
		// 本地连接
		DSN: "root:root@tcp(localhost:13316)/webook",
	},
	Redis: RedisConfig{
		Addr: "localhost:6379",
	},
}
