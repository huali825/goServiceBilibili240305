package main

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

	//第三种方式
	server := InitWebServer()
	_ = server.Run(":8080")

}
