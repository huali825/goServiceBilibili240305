package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler interface {

	// SetLoginToken 登录的时候生成长短 token，同时在这两个 token 里面都带上一个 ssid
	SetLoginToken(ctx *gin.Context, uid int64) error

	//SetJWTToken 短token
	SetJWTToken(ctx *gin.Context, uid int64, ssid string) error

	//CheckSession 功能是 检测session是否存在
	CheckSession(ctx *gin.Context, ssid string) error

	ExtractToken(ctx *gin.Context) string

	//ClearToken 删除token 退出登录的时候使用
	ClearToken(ctx *gin.Context) error
}

type RefreshClaims struct {
	Uid  int64
	Ssid string
	jwt.RegisteredClaims
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 声明你自己的要放进去 token 里面的数据
	Uid  int64
	Ssid string
	// 自己随便加
	UserAgent string
}
