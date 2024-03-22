package web

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
}

func NewUserHandler() *UserHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailRegexExp:    emailExp,
		passwordRegexExp: passwordExp,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) Profile(context *gin.Context) {
	context.String(http.StatusOK, "这是你的 Profile")
}

func (u *UserHandler) SignUp(context *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	if err := context.Bind(&req); err != nil {
		return
	}

	isEmail, err := u.emailRegexExp.MatchString(req.Email)
	if err != nil {
		context.String(http.StatusOK, "系统错误")
		return
	}
	if !isEmail {
		context.String(http.StatusOK, "邮箱格式不正确")
		return
	}
	if req.Password != req.ConfirmPassword {
		context.String(http.StatusOK, "两次密码输入不相同")
		return
	}

	isPassword, err2 := u.passwordRegexExp.MatchString(req.Password)
	if err2 != nil {
		context.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		context.String(http.StatusOK, "密码必须包含数字、特殊字符，并且长度不能小于 8 位")
		return
	}

	context.String(http.StatusOK, "注册成功了")

	fmt.Printf("%v", req)
}

func (u *UserHandler) Login(context *gin.Context) {
	context.String(http.StatusOK, "这是你的 login")
}

func (u *UserHandler) Edit(context *gin.Context) {
	context.String(http.StatusOK, "这是你的 edit")
}
