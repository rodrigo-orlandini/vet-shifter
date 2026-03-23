package controllers

import (
	"net/http"

	api "rodrigoorlandini/vet-shifter/internal/_shared/api"
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
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
//	@Failure		400		{object}	api.ApiErrorResponse	"Requisição inválida"
//	@Router			/auth/forgot-password [post]
func (c *ForgotPasswordController) Handle(ctx *gin.Context) {
	internalErr := &customerror.InternalServerError{}

	var body ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_REQUEST",
			"error": "Dados inválidos. Verifique os campos e tente novamente.",
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
		if _, ok := err.(*customerror.ServiceUnavailableError); ok {
			_ = ctx.Error(err)
			ctx.JSON(http.StatusServiceUnavailable, gin.H{
				"code":  "SERVICE_UNAVAILABLE",
				"error": err.Error(),
			})
			return
		}

		_ = ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":  "INTERNAL_ERROR",
			"error": internalErr.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "Se existir uma conta com este e-mail, você receberá um link para redefinir a senha.",
	})
}

// Swag references api.ApiErrorResponse in godoc comments.
var _ = api.ApiErrorResponse{}
