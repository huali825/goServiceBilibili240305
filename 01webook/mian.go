package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go20240218/01webook/internal/web"
	"strings"
	"time"
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
	server := initWebServer()
	u := web.NewUserHandler()
	u.RegisterRoutes(server)
	_ = server.Run(":8080")

}

func initWebServer() *gin.Engine {
	server := gin.Default()

	server.Use(func(context *gin.Context) {
		fmt.Println("tmh: first middleware")
	})
	server.Use(func(context *gin.Context) {
		fmt.Println("tmh: second middleware")
	})
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"*"},
		//AllowMethods: []string{"POST", "GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		//ExposeHeaders: []string{"x-jwt-token"},
		// 是否允许你带 cookie 之类的东西
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				// 你的开发环境
				return true
			}
			return strings.Contains(origin, "yourcompany.com")
		},
		MaxAge: 12 * time.Hour,
	}))
	return server
}
