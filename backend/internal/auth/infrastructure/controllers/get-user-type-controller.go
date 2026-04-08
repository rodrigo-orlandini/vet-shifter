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

type GetUserTypeResponse struct {
	UserType string `json:"user_type"`
}

type GetUserTypeController struct{}

func NewGetUserTypeController() *GetUserTypeController {
	return &GetUserTypeController{}
}

// GetUserTypeByEmail godoc
//
//	@Summary		Get user type by email
//	@Description	Checks company_owners and shift_veterinaries by email. Returns user_type for login redirect.
//	@Tags			auth
//	@Produce		json
//	@Param			email	query		string	true	"User email"
//	@Success		200		{object}	GetUserTypeResponse	"company_owner or shift_veterinary"
//	@Failure		400		{object}	api.ApiErrorResponse	"E-mail inválido"
//	@Failure		401		{object}	api.ApiErrorResponse	"Credenciais inválidas"
//	@Router			/auth/user-type [get]
func (c *GetUserTypeController) Handle(ctx *gin.Context) {
	internalErr := &customerror.InternalServerError{}

	emailStr := ctx.Query("email")
	if emailStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_REQUEST",
			"error": "Dados inválidos. Verifique os campos e tente novamente.",
		})
		return
	}

	email, err := sharedvalueobjects.NewEmail(emailStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_REQUEST",
			"error": err.Error(),
		})
		return
	}

	useCase := factories.NewGetUserTypeByEmailFactory()
	out, err := useCase.Execute(&usecases.GetUserTypeByEmailUseCaseInput{Email: *email})

	if err != nil {
		if _, ok := err.(*customerror.NotFoundError); ok {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":  "NOT_FOUND",
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

	ctx.JSON(http.StatusOK, GetUserTypeResponse{
		UserType: out.UserType.GetValue(),
	})
}

var _ = api.ApiErrorResponse{}
