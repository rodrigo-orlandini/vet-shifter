package controllers

import (
	"net/http"
	"time"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	usecases "rodrigoorlandini/vet-shifter/internal/veterinaries/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/factories"
	"rodrigoorlandini/vet-shifter/internal/veterinaries/infrastructure/mappers"

	"github.com/gin-gonic/gin"
)

type RegisterShiftVeterinaryRequest struct {
	FullName    string   `json:"full_name" binding:"required"`
	Cpf         string   `json:"cpf" binding:"required"`
	Email       string   `json:"email" binding:"required,email"`
	Phone       string   `json:"phone" binding:"required"`
	CrmvNumber  string   `json:"crmv_number" binding:"required"`
	CrmvState   string   `json:"crmv_state" binding:"required"`
	Specialties []string `json:"specialties" binding:"required"`
	Password    string   `json:"password" binding:"required"`
	ConsentLgpd bool     `json:"consent_lgpd" binding:"required"`
}

type RegisterShiftVeterinaryResponse struct {
	VeterinaryId string `json:"veterinary_id"`
}

type ErrorResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

type RegisterShiftVeterinaryController struct{}

func NewRegisterShiftVeterinaryController() *RegisterShiftVeterinaryController {
	return &RegisterShiftVeterinaryController{}
}

// RegisterShiftVeterinary godoc
//
//	@Summary		Register a new shift veterinary
//	@Description	Creates a veterinary account. Requires LGPD consent.
//	@Tags			veterinaries
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RegisterShiftVeterinaryRequest	true	"Veterinary data"
//	@Success		201		{object}	RegisterShiftVeterinaryResponse	"Created with veterinary_id"
//	@Failure		400		{object}	ErrorResponse	"Corpo da requisição inválido ou erro de validação"
//	@Failure		409		{object}	ErrorResponse	"CPF ou e-mail já cadastrados"
//	@Failure		500		{object}	ErrorResponse	"Erro interno do servidor"
//	@Router			/veterinaries [post]
func (c *RegisterShiftVeterinaryController) Handle(ctx *gin.Context) {
	internalErr := &customerror.InternalServerError{}

	var body RegisterShiftVeterinaryRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_REQUEST_BODY",
			"error": "Dados inválidos. Verifique os campos e tente novamente.",
		})
		return
	}

	var consentAt *time.Time
	if body.ConsentLgpd {
		t := time.Now()
		consentAt = &t
	}

	veterinary, err := mappers.ShiftVeterinaryFromHttp(mappers.ShiftVeterinaryFromHttpInput{
		Email:         body.Email,
		Phone:         body.Phone,
		Password:      body.Password,
		FullName:      body.FullName,
		Cpf:           body.Cpf,
		CrmvNumber:    body.CrmvNumber,
		CrmvState:     body.CrmvState,
		Specialties:   body.Specialties,
		ConsentLgpd:   body.ConsentLgpd,
		ConsentLgpdAt: consentAt,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_VETERINARY",
			"error": err.Error(),
		})
		return
	}

	useCase := factories.NewRegisterShiftVeterinaryFactory()
	out, err := useCase.Execute(&usecases.RegisterShiftVeterinaryUseCaseInput{Veterinary: *veterinary})
	if err != nil {
		switch err.(type) {
		case *customerror.InvalidValueObjectError:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  "INVALID_INPUT",
				"error": err.Error(),
			})
			return
		case *customerror.AlreadyExistsError:
			ctx.JSON(http.StatusConflict, gin.H{
				"code":  "ALREADY_EXISTS",
				"error": err.Error(),
			})
			return
		default:
			_ = ctx.Error(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":  "INTERNAL_SERVER_ERROR",
				"error": internalErr.Error(),
			})
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{"veterinary_id": out.VeterinaryId})
}
