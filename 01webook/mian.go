package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go20240218/01webook/internal/repository"
	"go20240218/01webook/internal/repository/dao"
	"go20240218/01webook/internal/service"
	"go20240218/01webook/internal/web"
	"go20240218/01webook/internal/web/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	db := initDB()
	uHandler := initDDD(db)
	server := initMiddleware()

	uHandler.RegisterRoutes(server)

	_ = server.Run(":8080")

}

func initMiddleware() *gin.Engine {
	server := gin.Default()

	server.Use(func(context *gin.Context) {
		fmt.Println("tmh: first middleware")
	})
	server.Use(func(context *gin.Context) {
		fmt.Println("tmh: second middleware")
	})
	//service.Use(cors.Default())
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

	//设置 cookie
	store := cookie.NewStore([]byte("secret"))
	server.Use(sessions.Sessions("mySession", store))

	//server.Use(func(context *gin.Context) {
	//
	//	if context.Request.URL.Path == "/users/login" ||
	//		context.Request.URL.Path == "/users/signup" {
	//		return
	//	}
	//	sess := sessions.Default(context)
	//	id := sess.Get("userId")
	//	if id == nil {
	//		context.AbortWithStatus(http.StatusUnauthorized)
	//		return
	//	}
	//})

	server.Use(middleware.NewLoginMiddlewareBuilder().
		IgnorePaths("/users/login").
		IgnorePaths("/users/signup").Build())

	return server
}

func initDDD(db *gorm.DB) *web.UserHandler {
	uDAO := dao.NewUserDAO(db)
	uRepo := repository.NewUserRepository(uDAO)
	uSvc := service.NewUserService(uRepo)
	uHandler := web.NewUserHandler(uSvc)
	return uHandler
}

func initDB() *gorm.DB {
	//初始化数据库
	//db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	if err != nil {
		fmt.Println("tmh: 数据库连接失败")
		panic(err)
	}

	err2 := dao.InitTable(db)
	if err2 != nil {
		fmt.Println("tmh: 数据库建表失败")
		panic(err2)
	}
	return db
}
