package main

import (
	"fmt"
)

func main() {

	//gin第1种方法
	//serverGin := gin.Default()
	//serverGin.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello gin server")
	//})
	//serverGin.Run("8081")

	//gin第2种方法
	//serverGin := gin.Default()
	//uHandler := &UserHandler{}
	//uHandler.RegisterRoutes(serverGin)
	//serverGin.Run("8081")

	server := NewServer("127.0.0.1", 8080)
	server.Start()
	fmt.Println("hello world go")
}
