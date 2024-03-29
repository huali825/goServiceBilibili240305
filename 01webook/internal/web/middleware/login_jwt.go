package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
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
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("95osj3fUD7fo0mlYdDbncXz4VD2igvf0"), nil
		})
		if err != nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//  err == nil
		if token == nil || !token.Valid {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
