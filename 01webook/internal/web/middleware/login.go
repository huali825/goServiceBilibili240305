package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		//if context.Request.URL.Path == "/users/login" ||
		//	context.Request.URL.Path == "/users/signup" {
		//	return
		//}
		for _, path := range l.paths {
			if context.Request.URL.Path == path {
				return
			}
		}
		sess := sessions.Default(context)
		id := sess.Get("userId")
		if id == nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
