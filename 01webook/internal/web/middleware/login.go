package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		gob.Register(time.Now())
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
			//没有登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		updateTime := sess.Get("update_time")
		sess.Set("userId", id)
		sess.Options(sessions.Options{
			MaxAge: 12,
		})

		timeNow := time.Now().UnixMilli()
		//刚登录 还没刷新
		if updateTime == nil {
			sess.Set("update_time", timeNow)
			sess.Save()
			return
		}
		updateTimeVal, ok := updateTime.(int64)
		if !ok {
			context.String(http.StatusInternalServerError, "有人搞我")
			return
		}
		if timeNow-updateTimeVal > 5*1000 {
			sess.Set("update_time", timeNow)
			sess.Save()
			return
		}
	}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}
