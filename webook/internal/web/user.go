package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-test/webook/internal/constants"
	"go-test/webook/internal/domain"
	"go-test/webook/internal/service"
	"go-test/webook/pkg"
	"net/http"
	"strconv"
	"time"
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
	serve.GET("/user", u.getUserInfo)
	serve.PUT("/register", u.addUser)
	serve.POST("/user", u.editUser)
	serve.DELETE("/user/:id", u.deleteUser)
	serve.POST("login", u.login)
	serve.POST("login_jwt", u.loginJwt)
}

func (u *UserHandler) addUser(ctx *gin.Context) {
	type UserReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
		NICKName        string `json:"nickName"`
		Profile         string `json:"profile"`
		BirthDay        string `json:"birthDay"`
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
		NickName: user.NICKName,
		Profile:  user.Profile,
		BirthDay: user.BirthDay,
	})
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}

	ctx.String(http.StatusOK, "注册成功！")
}

func (u *UserHandler) editUser(ctx *gin.Context) {
	type UserReq struct {
		Id       string `json:"id"`
		NickName string `json:"nickName"`
		Profile  string `json:"profile"`
		BirthDay string `json:"birthDay"`
	}

	var user UserReq
	if err := ctx.Bind(&user); err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	var newId int64
	if user.Id != "" {
		newId, _ = strconv.ParseInt(user.Id, 10, 64)
	} else {
		session := sessions.Default(ctx)
		sessionId := session.Get(constants.USER_ID).(int64)
		newId = sessionId
	}

	err := u.svc.UpdateUser(ctx, newId, domain.User{
		NickName: user.NickName,
		Profile:  user.Profile,
		BirthDay: user.BirthDay,
	})
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	ctx.String(http.StatusOK, "更新成功")
}

func (u *UserHandler) getUserInfo(ctx *gin.Context) {
	uc, _ := ctx.Get(constants.CLAIMS_KEY)
	userId := uc.(UserClaims).Id
	user, err := u.svc.GetUserInfo(ctx, userId)
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *UserHandler) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	err = u.svc.DeleteUser(ctx, intId)
	if err != nil {
		ctx.String(http.StatusOK, err.Error())
		return
	}
	ctx.String(http.StatusOK, "删除成功")
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
	session.Set(constants.USER_ID, user.Id)
	session.Options(sessions.Options{
		MaxAge: 10,
	})
	session.Save()
	ctx.String(http.StatusOK, "登录成功！")
}

func (u *UserHandler) loginJwt(ctx *gin.Context) {
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Id: user.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	})
	longTokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
	})
	shortToken, err := token.SignedString([]byte(constants.SHORT_TIME_JWT_KEY))
	longToken, err := longTokenClaims.SignedString([]byte(constants.LONG_TIME_JWT_KEY))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	pkg.Redis.Client.Set(ctx, constants.SHORT_TIME_JWT_KEY, shortToken, time.Hour)
	pkg.Redis.Client.Set(ctx, shortToken, longToken, time.Hour*24*30)
	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": shortToken,
		"msg":   "登录成功",
	})
}

type UserClaims struct {
	jwt.RegisteredClaims
	Id int64
}
