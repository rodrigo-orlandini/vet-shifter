package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	companiescontrollers "rodrigoorlandini/vet-shifter/internal/companies/infrastructure/controllers"
	"rodrigoorlandini/vet-shifter/test/integration"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginCompanyOwnerController_Handle(t *testing.T) {
	integration.PrepareDB(t)
	gin.SetMode(gin.TestMode)

	ctrl := NewLoginCompanyOwnerController()

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

	t.Run("happy path - returns 200 and access_token", func(t *testing.T) {
		integration.PrepareDB(t)
		email, pass := "owner@example.com", "password123"
		registerCompany(email, pass)
		body, _ := json.Marshal(map[string]interface{}{"email": email, "password": pass, "remember_me": false})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/auth/login/owner", bytes.NewReader(body))
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
		email := "owner2@example.com"
		registerCompany(email, "password123")
		body, _ := json.Marshal(map[string]interface{}{"email": email, "password": "wrong", "remember_me": false})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/auth/login/owner", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		ctrl.Handle(c)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "INVALID_CREDENTIALS", res["code"])
	})
}
