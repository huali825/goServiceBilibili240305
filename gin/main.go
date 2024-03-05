package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//静态路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK,
			" hello, go tmh first webservice")
	})

	r.GET("/users/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK,
			"参数路由, 参数为: "+name)
	})

	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
