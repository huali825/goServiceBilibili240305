package middleware

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	ijwt "go20240218/01webook/internal/web/jwt"
	"net/http"
)

type LoginJWTMiddlewareBuilder struct {
	paths []string
	ijwt.Handler
}

func NewLoginJWTMiddlewareBuilder(jwtHdl ijwt.Handler) *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{
		Handler: jwtHdl,
	}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		//gob.Register(time.Now())

		for _, path := range l.paths {
			if context.Request.URL.Path == path {
				return
			}
		}

		//使用长短token来登录校验
		tokenStr := l.ExtractToken(context)
		claims := &ijwt.UserClaims{}

		// ParseWithClaims 里面，一定要传入指针
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"), nil
		})
		if err != nil {
			// 没登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//claims.ExpiresAt.Time.Before(time.Now()) {
		//	// 过期了
		//}
		// err 为 nil，token 不为 nil
		if token == nil || !token.Valid || claims.Uid == 0 {
			// 没登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if claims.UserAgent != context.Request.UserAgent() {
			// 严重的安全问题
			// 你是要监控
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		err = l.CheckSession(context, claims.Ssid)
		if err != nil {
			// 要么 redis 有问题，要么已经退出登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//改造长短token之后就用不上这个了
		//now := time.Now()
		//if claims.ExpiresAt.Sub(now) < time.Second*50 {
		//	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * 3))
		//	tokenStr, err = token.SignedString([]byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"))
		//	if err != nil {
		//		//记录日志 jwt续约失败
		//	}
		//	//吧这个值返回给前端
		//	context.Header("x-jwt-token", tokenStr)
		//}
		context.Set("claims", claims)
	}
}
