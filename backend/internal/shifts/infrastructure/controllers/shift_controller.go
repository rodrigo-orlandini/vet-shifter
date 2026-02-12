package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"rodrigoorlandini/vet-shifter/internal/shifts/application/use_cases"
	"rodrigoorlandini/vet-shifter/internal/shifts/domain/entities"
	"rodrigoorlandini/vet-shifter/internal/shifts/infrastructure/factories"
)

func parseTime(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func parseInt32(s string) (int32, error) {
	n, err := strconv.ParseInt(s, 10, 32)
	return int32(n), err
}

type CreateShiftRequest struct {
	CompanyID         string `json:"company_id" binding:"required,uuid"`
	StartsAt          string `json:"starts_at" binding:"required"`
	EndsAt            string `json:"ends_at" binding:"required"`
	Type              string `json:"type" binding:"required,oneof=emergency consultation surgery"`
	OfferedValueCents int64  `json:"offered_value_cents" binding:"required,min=0"`
	Requirements      string `json:"requirements"`
	Description       string `json:"description"`
	Location          string `json:"location"`
}

type ShiftController struct{}

func NewShiftController() *ShiftController {
	return &ShiftController{}
}

func (c *ShiftController) Create(ctx *gin.Context) {
	var body CreateShiftRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_REQUEST_BODY", "error": err.Error()})
		return
	}
	startsAt, err := parseTime(body.StartsAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_DATETIME", "error": "starts_at invalid"})
		return
	}
	endsAt, err := parseTime(body.EndsAt)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_DATETIME", "error": "ends_at invalid"})
		return
	}
	shift, err := entities.NewShift(
		body.CompanyID,
		startsAt, endsAt,
		body.Type,
		body.OfferedValueCents,
		body.Requirements,
		body.Description,
		body.Location,
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "INVALID_SHIFT", "error": err.Error()})
		return
	}
	out, err := factories.NewCreateShiftFactory().Execute(&usecases.CreateShiftInput{Shift: *shift})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_SERVER_ERROR", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": out.Shift.Id})
}

func (c *ShiftController) List(ctx *gin.Context) {
	input := usecases.ListShiftsInput{
		Status:    ctx.Query("status"),
		Type:      ctx.Query("type"),
		CompanyID: ctx.Query("company_id"),
		FromDate:  ctx.Query("from_date"),
		ToDate:    ctx.Query("to_date"),
		Limit:     50,
		Offset:    0,
	}
	if l := ctx.Query("limit"); l != "" {
		if n, err := parseInt32(l); err == nil && n > 0 {
			input.Limit = n
		}
	}
	if o := ctx.Query("offset"); o != "" {
		if n, err := parseInt32(o); err == nil && n >= 0 {
			input.Offset = n
		}
	}
	out, err := factories.NewListShiftsFactory().Execute(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": "INTERNAL_SERVER_ERROR", "error": err.Error()})
		return
	}
	items := make([]map[string]interface{}, len(out.Shifts))
	for i, s := range out.Shifts {
		items[i] = map[string]interface{}{
			"id":                   s.Id,
			"company_id":           s.CompanyId,
			"starts_at":            s.StartsAt.Format("2006-01-02T15:04:05Z07:00"),
			"ends_at":              s.EndsAt.Format("2006-01-02T15:04:05Z07:00"),
			"type":                 s.Type,
			"offered_value_cents":  s.OfferedValueCents,
			"requirements":         s.Requirements,
			"description":          s.Description,
			"location":             s.Location,
			"status":               s.Status,
			"created_at":           s.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"shifts": items})
}
