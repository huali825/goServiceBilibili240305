package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go20240218/01webook/config"
	"go20240218/01webook/internal/repository"
	"go20240218/01webook/internal/repository/dao"
	"go20240218/01webook/internal/service"
	"go20240218/01webook/internal/web"
	"go20240218/01webook/internal/web/middleware"
	"go20240218/01webook/pkg/ginx/middlewares/ratelimit"
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
	//	str := fmt.Sprint("你好，你来了 webook hello！！！", config.Config.HS)
	//	ctx.String(http.StatusOK, str)
	//})
	//_ = server.Run(":8080")
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

	//需要在docker上面运行redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Config.Redis.Addr,
	})
	server.Use(ratelimit.NewBuilder(redisClient, time.Second, 100).Build())

	//service.Use(cors.Default())
	server.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"*"},
		//AllowMethods: []string{"POST", "GET"},

		//允许前端读取的东西
		AllowHeaders: []string{"Content-Type", "Authorization"},

		// 不用这个拿不到 jwt token
		ExposeHeaders: []string{"x-jwt-token"},

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
	//方法1 基于cookie的实现:
	//store := cookie.NewStore([]byte("secret"))

	//方法2 基于 memstore 的实现:
	//store := memstore.NewStore([]byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"),
	//	[]byte("0Pf2r0wZBpXVXlQNdpwCXN4ncnlnZSc3"))

	//方法3  redis 的实现
	//store, err := redis.NewStore(16, "tcp",
	//	"localhost:6379", "",
	//	[]byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"),
	//	[]byte("0Pf2r0wZBpXVXlQNdpwCXN4ncnlnZSc3"))
	//if err != nil {
	//	panic("连接redis失败!!!!!!!")
	//}

	//server.Use(sessions.Sessions("mySession", store))
	server.Use(middleware.NewLoginJWTMiddlewareBuilder().
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
	//db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/webook"))
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
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
