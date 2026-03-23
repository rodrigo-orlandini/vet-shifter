package controllers

import (
	"net/http"

	api "rodrigoorlandini/vet-shifter/internal/_shared/api"
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	autherrors "rodrigoorlandini/vet-shifter/internal/auth/application/custom-error"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/factories"

	"github.com/gin-gonic/gin"
)

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ResetPasswordController struct{}

func NewResetPasswordController() *ResetPasswordController {
	return &ResetPasswordController{}
}

// ResetPassword godoc
//
//	@Summary		Reset password with token
//	@Description	Consumes the token from email and sets new password. Token is invalidated.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		ResetPasswordRequest	true	"Token and new password"
//	@Success		200		{object}	map[string]string	"Success"
//	@Failure		400		{object}	api.ApiErrorResponse	"Token inválido ou senha fraca"
//	@Router			/auth/reset-password [post]
func (c *ResetPasswordController) Handle(ctx *gin.Context) {
	internalErr := &customerror.InternalServerError{}

	var body ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_REQUEST",
			"error": "Dados inválidos. Verifique os campos e tente novamente.",
		})
		return
	}

	useCase := factories.NewResetPasswordFactory()
	_, err := useCase.Execute(&usecases.ResetPasswordUseCaseInput{
		Token:       body.Token,
		NewPassword: body.NewPassword,
	})

	if err != nil {
		if _, ok := err.(*autherrors.InvalidResetTokenError); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  "INVALID_RESET_TOKEN",
				"error": err.Error(),
			})
			return
		}

		if _, ok := err.(*customerror.InvalidCredentialsError); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  "WEAK_PASSWORD",
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Senha atualizada com sucesso."})
}

// Swag references api.ApiErrorResponse in godoc comments.
var _ = api.ApiErrorResponse{}
