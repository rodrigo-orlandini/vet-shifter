package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "rodrigoorlandini/vet-shifter/internal/_shared/api"
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/factories"
	authmiddleware "rodrigoorlandini/vet-shifter/internal/auth/infrastructure/middleware"
)

type LoginVeterinaryRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}

type LoginVeterinaryResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

type LoginVeterinaryController struct{}

func NewLoginVeterinaryController() *LoginVeterinaryController {
	return &LoginVeterinaryController{}
}

// LoginVeterinary godoc
//
//	@Summary		Login as shift veterinary
//	@Description	Returns JWT for shift veterinaries. Token expiry 24h or 7d when remember_me is true.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		LoginVeterinaryRequest	true	"Credentials"
//	@Success		200		{object}	LoginVeterinaryResponse
//	@Failure		400		{object}	api.ApiErrorResponse	"Requisição inválida"
//	@Failure		401		{object}	api.ApiErrorResponse	"E-mail ou senha incorretos"
//	@Router			/auth/login/veterinary [post]
func (c *LoginVeterinaryController) Handle(ctx *gin.Context) {
	internalErr := &customerror.InternalServerError{}

	var body LoginVeterinaryRequest
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

	useCase := factories.NewLoginVeterinaryFactory()
	out, err := useCase.Execute(&usecases.LoginVeterinaryUseCaseInput{
		Email:      *email,
		Password:   body.Password,
		RememberMe: body.RememberMe,
	})

	if err != nil {
		if _, ok := err.(*customerror.InvalidCredentialsError); ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":  "INVALID_CREDENTIALS",
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

	authmiddleware.SetAccessTokenCookie(ctx, out.AccessToken, out.ExpiresAt)

	ctx.JSON(http.StatusOK, LoginVeterinaryResponse{
		AccessToken: out.AccessToken,
		ExpiresAt:   out.ExpiresAt,
	})
}

var _ = api.ApiErrorResponse{}
