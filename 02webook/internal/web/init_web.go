package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	registerUsersRoutes(server)
	return server
}

func registerUsersRoutes(server *gin.Engine) {
	u := &UserHandler{}
	server.POST("/users/signup", u.SignUp)
	// 这是 REST 风格
	//server.PUT("/user", func(context *gin.Context) {
	//
	//})

	server.POST("/users/login", u.Login)

	server.POST("/users/edit", u.Edit)
	// REST 风格
	//server.POST("/users/:id", func(context *gin.Context) {
	//
	//})

	server.GET("/users/profile", u.Profile)
	// REST 风格
	//server.GET("/users/:id", func(context *gin.Context) {
	//
	//})
}
