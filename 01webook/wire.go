//go:build wireinject

//让wire来注入这里的代码

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go20240218/01webook/internal/repository"
	"go20240218/01webook/internal/repository/cache"
	"go20240218/01webook/internal/repository/dao"
	"go20240218/01webook/internal/service"
	"go20240218/01webook/internal/web"
	ijwt "go20240218/01webook/internal/web/jwt"
	"go20240218/01webook/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 第三方依赖
		ioc.InitRedis, ioc.InitDB,

		// DAO 部分
		dao.NewUserDAO,

		// cache 部分
		cache.NewCodeCache, cache.NewUserCache,

		// repository 部分
		repository.NewCachedUserRepository,
		repository.NewCodeRepository,

		// Service 部分
		ioc.InitSMSService,
		service.NewUserService,
		service.NewCodeService,

		// handler 部分
		web.NewUserHandler,
		ioc.InitWebServerMiddleware,
		ijwt.NewRedisJWTHandler,
	)
	return gin.Default()
}
