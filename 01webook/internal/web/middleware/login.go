package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddlewareBuilder struct {
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.URL.Path == "/users/login" ||
			context.Request.URL.Path == "/users/signup" {
			return
		}
		sess := sessions.Default(context)
		id := sess.Get("userId")
		if id == nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
