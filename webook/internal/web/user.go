package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-test/webook/internal/domain"
	"go-test/webook/internal/service"
	"net/http"
)

type UserHandler struct {
	emailRegexExp    *regexp.Regexp
	passwordRegexExp *regexp.Regexp
	svc              *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegex          = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		passwordRegexStrong = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`
	)

	return &UserHandler{
		emailRegexExp:    regexp.MustCompile(emailRegex, regexp.None),
		passwordRegexExp: regexp.MustCompile(passwordRegexStrong, regexp.None),
		svc:              svc,
	}
}

func (u *UserHandler) RegisterUser(serve *gin.Engine) {
	serve.GET("/user/:id", u.getUserInfo)
	serve.PUT("/user", u.addUser)
	serve.POST("/user", u.editUser)
	serve.DELETE("/user/:id", u.deleteUser)
	serve.POST("login", u.login)
}

func (u *UserHandler) addUser(ctx *gin.Context) {
	type UserReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var user UserReq

	if err := ctx.Bind(&user); err != nil {
		println(err.Error())
		ctx.String(http.StatusOK, err.Error())
	}

	isEmail, err := u.emailRegexExp.MatchString(user.Email)
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	if !isEmail {
		ctx.String(http.StatusBadRequest, "邮箱格式错误！")
		return
	}
	isPassword, err := u.passwordRegexExp.MatchString(user.Password)
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	if !isPassword {
		ctx.String(http.StatusOK, "密码格式错误！")
		return
	}
	if user.Password != user.ConfirmPassword {
		ctx.String(http.StatusOK, "两次密码不一致！")
		return
	}

	err = u.svc.AddUser(ctx, domain.User{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}

	ctx.String(http.StatusOK, "注册成功！")
}

func (u *UserHandler) editUser(ctx *gin.Context) {
	println("editUser")
}

func (u *UserHandler) getUserInfo(ctx *gin.Context) {
	println("getUserInfo")
	ctx.String(http.StatusOK, "hello world")
}

func (u *UserHandler) deleteUser(ctx *gin.Context) {
	println("deleteUser")
}

func (u *UserHandler) login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	user, err := u.svc.Login(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}

	session := sessions.Default(ctx)
	session.Set("userId", user.Id)
	session.Save()
	ctx.String(http.StatusOK, "登录成功！")
}
