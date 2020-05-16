package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/fithrahfauzan/go-arch-example/internal/entity"
	"github.com/fithrahfauzan/go-arch-example/internal/usecase"
)

type AuthHTTP interface {
	LoginHandler(ctx *gin.Context)
	RegisterHandler(ctx *gin.Context)
	TokenValidationMiddleware() gin.HandlerFunc
}

type authHTTP struct {
	loginUseCase        usecase.UseCase
	registerUseCase     usecase.UseCase
	refreshTokenUseCase usecase.UseCase
}

func NewAuthHTTP(
	loginUseCase usecase.UseCase,
	registerUseCase usecase.UseCase,
	refreshTokenUseCase usecase.UseCase,
) AuthHTTP {
	return &authHTTP{loginUseCase, registerUseCase, refreshTokenUseCase}
}

func (a *authHTTP) LoginHandler(ctx *gin.Context) {
	var req entity.LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := a.loginUseCase.Execute(ctx, req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	WriteData(ctx, http.StatusOK, resp)
}

func (a *authHTTP) RegisterHandler(ctx *gin.Context) {
	var req entity.RegisterRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := a.registerUseCase.Execute(ctx, req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	WriteData(ctx, http.StatusOK, resp)
}

func (a *authHTTP) TokenValidationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
