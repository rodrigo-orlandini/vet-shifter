package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "rodrigoorlandini/vet-shifter/internal/_shared/api"
	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	sharedvalueobjects "rodrigoorlandini/vet-shifter/internal/_shared/value-objects"
	authmiddleware "rodrigoorlandini/vet-shifter/internal/auth/infrastructure/middleware"
	usecases "rodrigoorlandini/vet-shifter/internal/auth/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/auth/infrastructure/factories"
)

type LoginCompanyOwnerRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"remember_me"`
}

type LoginCompanyOwnerResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

type LoginCompanyOwnerController struct{}

func NewLoginCompanyOwnerController() *LoginCompanyOwnerController {
	return &LoginCompanyOwnerController{}
}

// LoginCompanyOwner godoc
//
//	@Summary		Login as company owner
//	@Description	Returns JWT for company owners. Token expiry 24h or 7d when remember_me is true.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		LoginCompanyOwnerRequest	true	"Credentials"
//	@Success		200		{object}	LoginCompanyOwnerResponse
//	@Failure		400		{object}	api.ApiErrorResponse	"Requisição inválida"
//	@Failure		401		{object}	api.ApiErrorResponse	"E-mail ou senha incorretos"
//	@Router			/auth/login/owner [post]
func (c *LoginCompanyOwnerController) Handle(ctx *gin.Context) {
	internalErr := &customerror.InternalServerError{}

	var body LoginCompanyOwnerRequest
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

	useCase := factories.NewLoginCompanyOwnerFactory()
	out, err := useCase.Execute(&usecases.LoginCompanyOwnerUseCaseInput{
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

	// Set HttpOnly cookie so the browser can authenticate subsequent requests.
	authmiddleware.SetAccessTokenCookie(ctx, out.AccessToken, out.ExpiresAt)

	ctx.JSON(http.StatusOK, LoginCompanyOwnerResponse{
		AccessToken: out.AccessToken,
		ExpiresAt:   out.ExpiresAt,
	})
}

// Swag references api.ApiErrorResponse in godoc comments.
var _ = api.ApiErrorResponse{}
