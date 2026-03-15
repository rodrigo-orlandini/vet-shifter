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

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/test/integration"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginVeterinaryController_Handle(t *testing.T) {
	integration.PrepareDB(t)
	gin.SetMode(gin.TestMode)

	ctrl := NewLoginVeterinaryController()

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

	t.Run("happy path - returns 200 and access_token", func(t *testing.T) {
		integration.PrepareDB(t)
		email, pass := "vet@example.com", "password123"
		insertShiftVeterinary(email, pass)
		body, _ := json.Marshal(map[string]interface{}{"email": email, "password": pass, "remember_me": false})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/auth/login/veterinary", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		ctrl.Handle(c)

		assert.Equal(t, http.StatusOK, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.NotEmpty(t, res["access_token"])
		assert.NotEmpty(t, res["expires_at"])
	})

	t.Run("main fail path - wrong password returns 401", func(t *testing.T) {
		integration.PrepareDB(t)
		email := "vet2@example.com"
		insertShiftVeterinary(email, "password123")
		body, _ := json.Marshal(map[string]interface{}{"email": email, "password": "wrong", "remember_me": false})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/auth/login/veterinary", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		ctrl.Handle(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "INVALID_CREDENTIALS", res["code"])
	})
}
