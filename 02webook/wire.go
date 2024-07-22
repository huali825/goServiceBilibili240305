//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go20240218/02webook/internal/repository"
	"go20240218/02webook/internal/repository/cache"
	"go20240218/02webook/internal/repository/dao"
	"go20240218/02webook/internal/service"
	"go20240218/02webook/internal/web"
	ijwt "go20240218/02webook/internal/web/jwt"
	"go20240218/02webook/ioc"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 最基础的第三方依赖
		ioc.InitDB, ioc.InitRedis,
		ioc.InitLogger,

		// 初始化 DAO
		dao.NewUserDAO,

		cache.NewUserCache,
		cache.NewCodeCache,

		repository.NewUserRepository,
		repository.NewCodeRepository,

		service.NewUserService,
		service.NewCodeService,
		// 直接基于内存实现
		ioc.InitSMSService,
		ioc.InitWechatService,

		web.NewUserHandler,
		web.NewOAuth2WechatHandler,
		ioc.NewWechatHandlerConfig,
		ijwt.NewRedisJWTHandler,
		// 你中间件呢？
		// 你注册路由呢？
		// 你这个地方没有用到前面的任何东西
		//gin.Default,

		ioc.InitWebServer,
		ioc.InitMiddlewares,
	)
	return new(gin.Engine)
}
