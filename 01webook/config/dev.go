//go:build !k8s

// asdsf go:build dev
// sdd go:build test
// dsf 34

// 没有k8s 这个编译标签

package config

var Config = config{
	DB: DBConfig{
		// 本地连接
		//DSN: "root:root@tcp(webook-mysql:30002)/webook",
		"root:root@tcp(localhost:30002)/webook", //连接 docker上运行的 mysql ： 可行√
		// "localhost:30002", // 不行 连接不上
	},
	Redis: RedisConfig{
		Addr: "localhost:30003",
	},
	HS: HelloString{
		str: "!k8s hello",
	},
}
