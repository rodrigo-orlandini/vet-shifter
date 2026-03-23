package controllers

import (
	"net/http"
	"time"

	customerror "rodrigoorlandini/vet-shifter/internal/_shared/custom-error"
	usecases "rodrigoorlandini/vet-shifter/internal/companies/application/use-cases"
	"rodrigoorlandini/vet-shifter/internal/companies/domain/entities"
	"rodrigoorlandini/vet-shifter/internal/companies/infrastructure/factories"
	"rodrigoorlandini/vet-shifter/internal/companies/infrastructure/mappers"

	"github.com/gin-gonic/gin"
)

type RegisterCompanyRequest struct {
	Cnpj        string `json:"cnpj" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
	OwnerName   string `json:"owner_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Street      string `json:"street"`
	Number      string `json:"number"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"zip_code"`
	ConsentLgpd bool   `json:"consent_lgpd" binding:"required"`
}

type RegisterCompanyResponse struct {
	CompanyId string `json:"company_id"`
}

type ErrorResponse struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

type RegisterCompanyController struct{}

func NewRegisterCompanyController() *RegisterCompanyController {
	return &RegisterCompanyController{}
}

// RegisterCompany godoc
//
//	@Summary		Register a new company with owner
//	@Description	Creates a company and its owner account. Requires LGPD consent.
//	@Tags			companies
//	@Accept			json
//	@Produce		json
//	@Param			body	body		RegisterCompanyRequest	true	"Company and owner data"
//	@Success		201		{object}	RegisterCompanyResponse	"Created with company_id"
//	@Failure		400		{object}	ErrorResponse	"Corpo da requisição inválido ou erro de validação"
//	@Failure		409		{object}	ErrorResponse	"CNPJ ou e-mail já cadastrados"
//	@Failure		500		{object}	ErrorResponse	"Erro interno do servidor"
//	@Router			/companies [post]
func (c *RegisterCompanyController) Handle(ctx *gin.Context) {
	internalErr := &customerror.InternalServerError{}

	var body RegisterCompanyRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_REQUEST_BODY",
			"error": "Dados inválidos. Verifique os campos e tente novamente.",
		})
		return
	}

	company, err := mappers.CompanyFromHttp(body.Cnpj, body.CompanyName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_COMPANY",
			"error": err.Error(),
		})
		return
	}

	var consentAt *time.Time
	if body.ConsentLgpd {
		t := time.Now()
		consentAt = &t
	}

	companyOwner, err := mappers.CompanyOwnerFromHttp(mappers.CompanyOwnerFromHttpInput{
		Email:         body.Email,
		Phone:         body.Phone,
		Password:      body.Password,
		CompanyId:     company.Id,
		ConsentLgpdAt: consentAt,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  "INVALID_OWNER",
			"error": err.Error(),
		})
		return
	}

	var address *entities.Address
	if body.Street != "" || body.City != "" || body.ZipCode != "" {
		addr, addrErr := entities.NewAddress(
			company.Id, body.Street, body.Number, body.City, body.State, body.ZipCode,
		)
		if addrErr != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code":  "INVALID_ADDRESS",
				"error": addrErr.Error(),
			})
			return
		}

		address = addr
	}

	input := usecases.RegisterCompanyUseCaseInput{
		Company:      *company,
		CompanyOwner: *companyOwner,
		Address:      address,
	}

	useCase := factories.NewRegisterCompanyFactory()
	out, err := useCase.Execute(&input)
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

	ctx.JSON(http.StatusCreated, gin.H{"company_id": out.CompanyId})
}
