package web

import (
	"errors"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"go20240218/01webook/internal/domain"
	"go20240218/01webook/internal/service"
	"net/http"
	"time"
)

//var (
//	ErrUserDuplicateEmail    = service.ErrUserDuplicateEmail
//	ErrInvalidUserOrPassword = service.ErrInvalidUserOrPassword
//)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
	svc              *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,32}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailRegexExp:    emailExp,
		passwordRegexExp: passwordExp,
		svc:              svc,
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

// SignUp 注册
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

	isPassword, err := u.passwordRegexExp.MatchString(req.Password)
	if err != nil {
		context.String(http.StatusOK, "系统错误")
		return
	}
	if !isPassword {
		context.String(http.StatusOK, "密码必须包含数字、特殊字符，并且长度不能小于 8 位")
		return
	}

	fmt.Printf("%v", req)
	//往下进行业务
	err = u.svc.Signup(context, domain.User{
		Id:       0,
		Email:    req.Email,
		Password: req.Password,
		Ctime:    time.Time{},
		Dtime:    time.Time{},
	})
	if errors.Is(err, service.ErrUserDuplicateEmail) {
		context.String(http.StatusOK, "数据库注册 邮箱冲突")
		return
	}
	if err != nil {
		context.String(http.StatusOK, "数据库注册 系统错误")
		return
	}

	context.String(http.StatusOK, "注册成功了")
}

// Login 登录
func (u *UserHandler) Login(context *gin.Context) {
	type LoginReq struct {
		Email    string
		Password string
	}
	var loginReq LoginReq
	if err := context.Bind(&loginReq); err != nil {
		return
	}

	user, err := u.svc.Login(context, domain.User{
		Email:    loginReq.Email,
		Password: loginReq.Password,
	})
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		context.String(http.StatusOK, "用户名或密码不对")
		return
	}
	if err != nil {
		context.String(http.StatusOK, "系统错误")
		return
	}

	fmt.Println(user) //temp
	//context.String(http.StatusOK, "这是你的 login")
}

// Edit 编辑
func (u *UserHandler) Edit(context *gin.Context) {
	context.String(http.StatusOK, "这是你的 edit")
}
