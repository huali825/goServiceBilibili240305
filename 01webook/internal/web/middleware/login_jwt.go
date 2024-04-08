package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"go20240218/01webook/internal/web"
	"net/http"
	"strings"
	"time"
)

type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
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

		//我现在 使用 JWT 来校验
		tokenHeader := context.GetHeader("Authorization")
		if tokenHeader == "" {
			//如果没有 就是没有登录   error 401
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			//error 401
			context.AbortWithStatus(http.StatusUnauthorized)
			// 有人搞鬼
			return
		}

		tokenStr := segs[1]
		claims := &web.UserClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"), nil
		})
		if err != nil {
			//没登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//  err == nil
		if token == nil || !token.Valid || claims.Uid == 0 {
			// jwt 你妹登录
			fmt.Println("jwt login : 你妹登录!!")
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.UserAgent != context.Request.UserAgent() {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		now := time.Now()
		if claims.ExpiresAt.Sub(now) < time.Second*50 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * 3))
			tokenStr, err = token.SignedString([]byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"))
			if err != nil {
				//记录日志 jwt续约失败
			}
			//吧这个值返回给前端
			context.Header("x-jwt-token", tokenStr)
		}
		context.Set("claims", claims)
	}
}
