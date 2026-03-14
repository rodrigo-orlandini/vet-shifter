package controllers

import (
	"net/http"

	_ "rodrigoorlandini/vet-shifter/internal/_shared/api"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/factories"

	"github.com/gin-gonic/gin"
)

type ForgotPasswordController struct{}

func NewForgotPasswordController() *ForgotPasswordController {
	return &ForgotPasswordController{}
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ForgotPassword godoc
//
//	@Summary		Request password reset
//	@Description	Sends reset link to email if account exists. Always returns 202.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		ForgotPasswordRequest	true	"Email"
//	@Success		202		{object}	map[string]string	"Accepted"
//	@Failure		400		{object}	api.ApiErrorResponse	"Invalid request"
//	@Router			/auth/forgot-password [post]
func (c *ForgotPasswordController) Handle(ctx *gin.Context) {
	var body ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_REQUEST",
			"error": err.Error(),
		})
		return
	}

	email, err := sharedvalueobjects.NewEmail(body.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_REQUEST",
			"error": err.Error(),
		})
		return
	}

	useCase := factories.NewRequestPasswordResetFactory()
	_, err = useCase.Execute(&usecases.RequestPasswordResetUseCaseInput{Email: *email})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  "INTERNAL_ERROR",
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "If an account exists with this email, you will receive a password reset link.",
	})
}
