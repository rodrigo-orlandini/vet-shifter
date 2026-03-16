package controllers

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	companiescontrollers "rodrigoorlandini/vet-shifter/internal/companies/infrastructure/controllers"
	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/test/integration"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetUserTypeController_Handle(t *testing.T) {
	integration.PrepareDB(t)
	gin.SetMode(gin.TestMode)

	ctrl := NewGetUserTypeController()

	registerCompany := func(email, password string) {
		t.Helper()
		body, _ := json.Marshal(map[string]interface{}{
			"cnpj": "11222333000181", "company_name": "Clinic", "owner_name": "Owner",
			"email": email, "phone": "11987654321", "password": password, "consent_lgpd": true,
		})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")
		reg := companiescontrollers.NewRegisterCompanyController()
		reg.Handle(c)
		require.Equal(t, http.StatusCreated, w.Code)
	}

	insertShiftVeterinary := func(email, password string) {
		t.Helper()
		conn := database.GetPostgresConnection()
		q := database.NewQueries(conn)
		ctx := context.Background()
		_, err := q.CreateShiftVeterinary(ctx, queries.CreateShiftVeterinaryParams{
			ID:             uuid.New(),
			Email:          email,
			Phone:          "11987654321",
			Password:       utils.Argon2Hash(password),
			FullName:       "Vet Name",
			Cpf:            "12345678901",
			CrmvNumber:     "12345",
			CrmvState:      "SP",
			Specialties:    []queries.VeterinarySpecialty{queries.VeterinarySpecialtyGeneralPractice},
			ApprovalStatus: queries.AccountStatusComplete,
			ConsentLgpdAt:  sql.NullTime{Valid: true, Time: time.Now()},
		})
		require.NoError(t, err)
	}

	t.Run("happy path - returns 200 and company_owner when email is owner", func(t *testing.T) {
		integration.PrepareDB(t)
		email := "owner@example.com"
		registerCompany(email, "password123")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/auth/user-type?email="+email, nil)

		ctrl.Handle(c)

		assert.Equal(t, http.StatusOK, w.Code)
		var res GetUserTypeResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "company_owner", res.UserType)
	})

	t.Run("happy path - returns 200 and shift_veterinary when email is veterinary", func(t *testing.T) {
		integration.PrepareDB(t)
		email := "vet@example.com"
		insertShiftVeterinary(email, "password123")

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/auth/user-type?email="+email, nil)

		ctrl.Handle(c)

		assert.Equal(t, http.StatusOK, w.Code)
		var res GetUserTypeResponse
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "shift_veterinary", res.UserType)
	})

	t.Run("returns 404 when email not found in any table", func(t *testing.T) {
		integration.PrepareDB(t)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodGet, "/auth/user-type?email=nobody@example.com", nil)

		ctrl.Handle(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "NOT_FOUND", res["code"])
		assert.Contains(t, res["error"], "no account found")
	})
}
