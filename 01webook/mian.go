package main

import (
	"github.com/spf13/viper"
)

//程序启动的入口

func main() {
	//初步使用
	//server := gin.Default()
	//server.GET("/hello", func(ctx *gin.Context) {
	//	str := fmt.Sprint("你好，你来了 webook hello！！！", config.Config.HS)
	//	ctx.String(http.StatusOK, str)
	//})
	//_ = server.Run(":8080")
	//第二种方式 不常用 实例看看 在init_web.go 里面实现的
	//server := web.RegisterRoutes()
	//server.Run(":8080")

	initViper()

	//第三种方式 wire的方式
	server := InitWebServer()
	_ = server.Run(":8080")

}

func initViper() {
	viper.SetDefault("db.mysql.dsn",
		"root:root@tcp(localhost:13316)/webook")

	// 配置文件的名字，但是不包含文件扩展名
	// 不包含 .go, .yaml 之类的后缀
	viper.SetConfigName("dev")
	// 告诉 viper 我的配置用的是 yaml 格式
	// 现实中，有很多格式，JSON，XML，YAML，TOML，ini
	viper.SetConfigType("yaml")
	// 当前工作目录下的 config 子目录
	viper.AddConfigPath("./config")
	//viper.AddConfigPath("/tmp/config")
	//viper.AddConfigPath("/etc/webook")
	// 读取配置到 viper 里面，或者你可以理解为加载到内存里面
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//你可以有多个viper的实例
	//otherViper := viper.New()
	//otherViper.SetConfigName("myjson")
	//otherViper.AddConfigPath("./config")
	//otherViper.SetConfigType("json")
}
