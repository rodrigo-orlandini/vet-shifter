package controllers

import (
	"net/http"
	"time"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	usecases "rodrigoorlandini/vet-shifter/internal/shift_vets/application/use_cases"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/infrastructure/factories"
	"rodrigoorlandini/vet-shifter/internal/shift_vets/infrastructure/mappers"

	"github.com/gin-gonic/gin"
)

type RegisterShiftVetRequest struct {
	Email         string   `json:"email" binding:"required,email"`
	Phone         string   `json:"phone" binding:"required"`
	Password      string   `json:"password" binding:"required"`
	FullName      string   `json:"full_name" binding:"required"`
	Cpf           string   `json:"cpf" binding:"required"`
	CrmvNumber    string   `json:"crmv_number" binding:"required"`
	CrmvState     string   `json:"crmv_state" binding:"required,len=2"`
	Specialties   []string `json:"specialties"`
	ConsentLgpd   bool     `json:"consent_lgpd" binding:"required"`
}

type RegisterShiftVetController struct{}

func NewRegisterShiftVetController() *RegisterShiftVetController {
	return &RegisterShiftVetController{}
}

func (c *RegisterShiftVetController) Handle(ctx *gin.Context) {
	var body RegisterShiftVetRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_REQUEST_BODY", "error": err.Error()})
		return
	}
	var consentAt *time.Time
	if body.ConsentLgpd {
		t := time.Now()
		consentAt = &t
	}
	vet, err := mappers.ShiftVetFromHttp(mappers.ShiftVetFromHttpInput{
		Email:         body.Email,
		Phone:         body.Phone,
		Password:      body.Password,
		FullName:      body.FullName,
		Cpf:           body.Cpf,
		CrmvNumber:    body.CrmvNumber,
		CrmvState:     body.CrmvState,
		Specialties:   body.Specialties,
		ConsentLgpdAt: consentAt,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_SHIFT_VET", "error": err.Error()})
		return
	}
	out, err := factories.NewRegisterShiftVetFactory().Execute(&usecases.RegisterShiftVetInput{Vet: *vet})
	if err != nil {
		switch e := err.(type) {
		case *customerror.AlreadyExistsError:
			ctx.JSON(http.StatusConflict, gin.H{"code": "ALREADY_EXISTS", "error": e.Error()})
			return
		case *customerror.RepositoryError:
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": "REPOSITORY_ERROR", "error": e.Error()})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_SERVER_ERROR", "error": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": out.Vet.Id})
}
