package web

import (
	"github.com/gin-gonic/gin"
	"go20240218/02webook/internal/domain"
	"go20240218/02webook/internal/service"
	"go20240218/02webook/pkg/logger"
	"net/http"
)

var _ handler = (*ArticleHandler)(nil)

type ArticleHandler struct {
	svc service.ArticleService
	l   logger.LoggerV1
}

func NewArticleHandler(
	svc service.ArticleService,
	l logger.LoggerV1) *ArticleHandler {
	return &ArticleHandler{
		svc: svc,
		l:   l,
	}
}

func (h *ArticleHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/articles")

	//编辑 也可以是新建 也可以修改
	ug.POST("/edit", h.Edit)
	//ug.POST("/signup", h.SignUp)
}

func (h *ArticleHandler) Edit(ctx *gin.Context) {
	type Req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}

	//检测输入
	//这里跳过

	//调用 service
	id, err := h.svc.Save(ctx, domain.Article{
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
			//Data: nil,
		})
		h.l.Error("保存帖子失败", logger.Error(err))
		return
	}
	ctx.JSON(http.StatusOK, Result{
		Msg:  "OK",
		Data: id,
	})

}
