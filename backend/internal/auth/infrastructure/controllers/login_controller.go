package controllers

import (
	"net/http"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use_cases"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/factories"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginController struct{}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (c *LoginController) Handle(ctx *gin.Context) {
	var body LoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_REQUEST_BODY", "error": err.Error()})
		return
	}
	out, err := factories.NewLoginFactory().Execute(&usecases.LoginInput{Email: body.Email, Password: body.Password})
	if err != nil {
		if _, ok := err.(*customerror.InvalidCredentialsError); ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": "INVALID_CREDENTIALS", "error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_SERVER_ERROR", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": out.Token, "role": out.Role, "sub": out.Sub})
}
