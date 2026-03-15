package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"rodrigoorlandini/vet-shifter/test/integration"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterCompanyController_Handle(t *testing.T) {
	integration.PrepareDB(t)
	gin.SetMode(gin.TestMode)

	ctrl := NewRegisterCompanyController()

	happyBody := map[string]interface{}{
		"cnpj":         "11222333000181",
		"company_name": "Test Vet Clinic",
		"owner_name":   "Owner Name",
		"email":        "owner@example.com",
		"phone":        "11987654321",
		"password":     "password123",
		"consent_lgpd": true,
	}

	t.Run("happy path - creates company and returns 201", func(t *testing.T) {
		integration.PrepareDB(t)
		body, _ := json.Marshal(happyBody)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		ctrl.Handle(c)

		assert.Equal(t, http.StatusCreated, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Contains(t, res, "company_id")
		assert.NotEmpty(t, res["company_id"])
	})

	t.Run("main fail path - duplicate CNPJ returns 409", func(t *testing.T) {
		integration.PrepareDB(t)
		body, _ := json.Marshal(happyBody)
		w1 := httptest.NewRecorder()
		c1, _ := gin.CreateTestContext(w1)
		c1.Request = httptest.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		c1.Request.Header.Set("Content-Type", "application/json")
		ctrl.Handle(c1)
		require.Equal(t, http.StatusCreated, w1.Code)

		w2 := httptest.NewRecorder()
		c2, _ := gin.CreateTestContext(w2)
		c2.Request = httptest.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		c2.Request.Header.Set("Content-Type", "application/json")
		ctrl.Handle(c2)

		assert.Equal(t, http.StatusConflict, w2.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w2.Body.Bytes(), &res))
		assert.Equal(t, "ALREADY_EXISTS", res["code"])
	})
}
