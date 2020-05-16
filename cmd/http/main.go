package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"

	"github.com/fithrahfauzan/go-arch-example/internal/config"
	"github.com/fithrahfauzan/go-arch-example/internal/data"
	"github.com/fithrahfauzan/go-arch-example/internal/presentation/http"
	"github.com/fithrahfauzan/go-arch-example/internal/repository"
	"github.com/fithrahfauzan/go-arch-example/internal/usecase"

	authusecase "github.com/fithrahfauzan/go-arch-example/internal/usecase/auth"
)

var (
	sqlDB       *sql.DB
	redisClient *redis.Client
	mainCfg     config.MainConfig
)

var (
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
)

var (
	loginUseCase        usecase.UseCase
	registerUseCase     usecase.UseCase
	refreshTokenUseCase usecase.UseCase
)

var (
	authHTTP http.AuthHTTP
)

func main() {
	initAppModule()
	initRepo()
	initUseCase()
	initHTTPPresentation()

	run()
}

func initAppModule() {
	sqlDB = &sql.DB{}
	redisClient = &redis.Client{}
	mainCfg = config.NewMainConfig()
}

func initRepo() {
	userRepo = data.NewUserRepository(sqlDB, redisClient)
	tokenRepo = data.NewtokenRepository(sqlDB, redisClient, mainCfg.Repository)
}

func initUseCase() {
	loginUseCase = authusecase.NewLoginUseCase(userRepo, tokenRepo)
	registerUseCase = authusecase.NewRegisterUseCase(userRepo, tokenRepo)
	refreshTokenUseCase = authusecase.NewRefreshTokenUseCase(tokenRepo)
}

func initHTTPPresentation() {
	authHTTP = http.NewAuthHTTP(loginUseCase, registerUseCase, refreshTokenUseCase)
}

func run() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/auth/login", authHTTP.LoginHandler)
	r.POST("/auth/register", authHTTP.RegisterHandler)
	r.Use(authHTTP.TokenValidationMiddleware())

	r.Run(":3000")
}
