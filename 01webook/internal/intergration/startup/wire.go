//go:build wireinject

package startup

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

var thirdProvider = wire.NewSet(InitRedis, InitTestDB, InitLog)
var userSvcProvider = wire.NewSet(
	dao.NewUserDAO,
	cache.NewUserCache,
	repository.NewUserRepository,
	service.NewUserService)

func InitWebServer() *gin.Engine {
	wire.Build(
		thirdProvider,
		userSvcProvider,
		//articlSvcProvider,
		cache.NewCodeCache,
		dao.NewGORMArticleDAO,
		repository.NewCodeRepository,
		repository.NewArticleRepository,
		// service 部分
		// 集成测试我们显式指定使用内存实现
		ioc.InitSMSService,

		// 指定啥也不干的 wechat service
		InitPhantomWechatService,
		service.NewCodeService,
		service.NewArticleService,
		// handler 部分
		web.NewUserHandler,
		web.NewOAuth2WechatHandler,
		web.NewArticleHandler,
		InitWechatHandlerConfig,
		ijwt.NewRedisJWTHandler,

		// gin 的中间件
		ioc.InitMiddlewares,

		// Web 服务器
		ioc.InitWebServer,
	)
	// 随便返回一个
	return gin.Default()
}

func InitArticleHandler() *web.ArticleHandler {
	wire.Build(thirdProvider,
		dao.NewGORMArticleDAO,
		service.NewArticleService,
		web.NewArticleHandler,
		repository.NewArticleRepository,
	)
	return &web.ArticleHandler{}
}

func InitUserSvc() service.UserService {
	wire.Build(thirdProvider, userSvcProvider)
	return service.NewUserService(nil, nil)
}

func InitJwtHdl() ijwt.Handler {
	wire.Build(thirdProvider, ijwt.NewRedisJWTHandler)
	return ijwt.NewRedisJWTHandler(nil)
}
