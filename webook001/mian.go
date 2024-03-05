package main

import (
	"github.com/gin-gonic/gin"
	"go20240218/webook001/internal/web"
)

//程序启动的入口

func main() {
	//初步使用
	//server := gin.Default()
	//server.GET("/hello", func(ctx *gin.Context) {
	//	ctx.String(http.StatusOK, "你好，你来了")
	//})
	//server.Run(":8080")

	//第二种方式 不常用 实例看看 在init_web.go 里面实现的
	//server := web.RegisterRoutes()
	//server.Run(":8080")

	//第三种方式
	server := gin.Default()
	u := &web.UserHandler{}
	u.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {

	}
}
