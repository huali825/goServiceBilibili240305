package main

import (
	"github.com/gin-gonic/gin"
	"go20240218/01webook/internal/repository"
	"go20240218/01webook/internal/repository/cache"
	"go20240218/01webook/internal/repository/dao"
	"go20240218/01webook/internal/service"
	"go20240218/01webook/internal/web"
	"go20240218/01webook/ioc"
)

func InitWebServer() *gin.Engine {
	db := ioc.InitDB()
	rdb := ioc.InitRedis()
	smsService := ioc.InitSMSService()

	uDAO := dao.NewUserDAO(db)
	uRepo := repository.NewUserRepository(uDAO)
	uSvc := service.NewUserService(uRepo)

	codeCache := cache.NewCodeCache(rdb)
	codeRep := repository.NewCodeRepository(codeCache)
	codeSvc := service.NewCodeService(codeRep, smsService)
	uHandler := web.NewUserHandler(uSvc, codeSvc)

	server := initMiddleware(rdb)

	uHandler.RegisterRoutes(server)

	return server
}

//func initDDD(db *gorm.DB) *web.UserHandler {
//	uDAO := dao.NewUserDAO(db)
//	uRepo := repository.NewUserRepository(uDAO)
//	uSvc := service.NewUserService(uRepo)
//
//	codeCache := cache.NewCodeCache(rbd)
//	codeRep := repository.NewCodeRepository()
//	codeSvc := service.NewCodeService(codeRep)
//	uHandler := web.NewUserHandler(uSvc, codeSvc)
//	return uHandler
//}
