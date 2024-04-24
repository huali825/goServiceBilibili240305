package web

import (
	"errors"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"go20240218/01webook/internal/domain"
	"go20240218/01webook/internal/service"
	"net/http"
	"time"
)

//var (
//	ErrUserDuplicateEmail    = service.ErrUserDuplicateEmail
//	ErrInvalidUserOrPassword = service.ErrInvalidUserOrPassword
//)

const (
	emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,32}$`
	bizLogin             = "login"
)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
	svc              service.UserService
	codeSvc          service.CodeService
}

func NewUserHandler(svc service.UserService, codeSvc service.CodeService) *UserHandler {

	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailRegexExp:    emailExp,
		passwordRegexExp: passwordExp,
		svc:              svc,
		codeSvc:          codeSvc,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.ProfileJWT)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.LoginJWT)
	ug.POST("/edit", u.Edit)

	// 手机验证码登录相关功能
	ug.POST("/login_sms/code/send", u.SendSMSLoginCode)
	ug.POST("/login_sms", u.LoginSMS)
}

func (u *UserHandler) Profile(context *gin.Context) {
	context.String(http.StatusOK, "这是你的 Profile")
}

func (u *UserHandler) ProfileJWT(context *gin.Context) {
	c, _ := context.Get("claims")
	//if !ok {
	//	context.String(http.StatusOK, "系统错误")
	//	return
	//}
	claims, ok := c.(*UserClaims)
	if !ok {
		context.String(http.StatusOK, "系统错误")
		return
	}

	fmt.Println(claims.Uid)
	context.String(http.StatusOK, "jwt profile")
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

	fmt.Println("web里面: ", req)
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
	return
}

// LoginJWT 登录JWT操作
func (u *UserHandler) LoginJWT(context *gin.Context) {
	type LoginReq struct {
		Email    string
		Password string
	}
	var loginReq LoginReq
	if err := context.Bind(&loginReq); err != nil {
		return
	}

	user, err := u.svc.Login(context, loginReq.Email, loginReq.Password)
	if errors.Is(err, service.ErrInvalidUserOrPassword) {
		context.String(http.StatusOK, "用户名或密码不对")
		return
	}
	if err != nil {
		context.String(http.StatusOK, "系统错误")
		return
	}

	//===下面设置登录态===//
	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 10)),
		},
		Uid: user.Id,

		UserAgent: context.Request.UserAgent(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	tokenStr, err := token.SignedString(JWTKey)
	if err != nil {
		context.String(http.StatusInternalServerError, "系统错误")
		return
	}
	fmt.Println("登陆成功的token str : ", tokenStr)

	context.Header("x-jwt-token", tokenStr)
	fmt.Println(user)
	context.String(http.StatusOK, "登录成功")
	return
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

	user, err := u.svc.Login(context, loginReq.Email, loginReq.Password)
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

	//登录成功 设置session
	sess := sessions.Default(context)
	sess.Set("userId", user.Id)
	sess.Options(sessions.Options{
		//Secure:   true,
		//HttpOnly: true,
		MaxAge: 12,
	})
	_ = sess.Save()

	context.String(http.StatusOK, "登录成功")
	return
}

// Edit 编辑
func (u *UserHandler) Edit(context *gin.Context) {
	context.String(http.StatusOK, "这是你的 edit")
}

func (u *UserHandler) SendSMSLoginCode(context *gin.Context) {
	fmt.Println("这里是 SendSMSLoginCode func")
	//context.String(http.StatusOK, "这里是 SendSMSLoginCode func")
	//return
	type Req struct {
		Phone string `json:"phone"`
	}
	var req Req
	if err := context.Bind(&req); err != nil {
		//context.String(http.StatusOK, err.Error())
		return
	}

	if req.Phone == "" {
		context.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "请输入手机号码",
		})
		return
	}
	err := u.codeSvc.Send(context, bizLogin, req.Phone)
	switch {
	case err == nil:
		context.JSON(http.StatusOK, Result{
			Msg: "发送成功11",
		})
	case errors.Is(err, service.ErrCodeSendTooMany):
		context.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "短信发送太频繁，请稍后再试",
		})
	default:
		context.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		// 补日志的
	}
}

func (u *UserHandler) LoginSMS(context *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	if err := context.Bind(&req); err != nil {
		return
	}

	ok, err := u.codeSvc.Verify(context, bizLogin, req.Phone, req.Code)
	if err != nil {
		context.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统异常",
		})
		return
	}
	if !ok {
		context.JSON(http.StatusOK, Result{
			Code: 4,
			Msg:  "验证码不对, 请重新输入",
		})
		return
	}

	duser, err := u.svc.FindOrCreate(context, req.Phone)
	if err != nil {
		context.JSON(http.StatusOK, Result{
			Code: 5,
			Msg:  "系统错误",
		})
		return
	}
	u.setJWTToken(context, duser.Id)
	context.JSON(http.StatusOK, Result{
		Msg: "登录成功",
	})
}

func (u *UserHandler) setJWTToken(ctx *gin.Context, uid int64) {
	uc := UserClaims{
		Uid:       uid,
		UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			// 1 分钟过期
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, uc)
	tokenStr, err := token.SignedString(JWTKey)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
	}
	ctx.Header("x-jwt-token", tokenStr)
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       int64
	UserAgent string
}

var JWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")
