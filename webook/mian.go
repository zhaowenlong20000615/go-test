package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-test/webook/internal/repository"
	"go-test/webook/internal/repository/dao"
	"go-test/webook/internal/service"
	"go-test/webook/internal/web"
	"go-test/webook/internal/web/middleware"
	"go-test/webook/pkg"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

func main() {
	db := initDB()
	initRedisDB()
	server := initWebServer()
	initUser(server, db)
	server.Run(":8080")
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13306)/webook"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTable(db)
	if err != nil {
		panic(err)
	}
	return db
}

func initRedisDB() {
	pkg.ConnectRedis()
}

func initWebServer() *gin.Engine {
	serve := web.RegisterRouters()

	serve.Use(cors.New(cors.Config{
		//AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "x-jwt-token"},
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "https://github.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	//store := cookie.NewStore([]byte("secret"))
	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	//serve.Use(sessions.Sessions("webook", store))
	//serve.Use(middleware.NewLoginMiddlewareBuilder().IgnorePaths("/login", "/register").Build())
	serve.Use(middleware.NewLoginJwtMiddlewareBuilder().IgnorePaths("/login_jwt", "/register").Build())
	return serve
}

func initUser(server *gin.Engine, db *gorm.DB) {
	d := dao.NewUserDAO(db)
	r := repository.NewUserRepository(d)
	s := service.NewUserService(r)
	h := web.NewUserHandler(s)
	h.RegisterUser(server)
}
