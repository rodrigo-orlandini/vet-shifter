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

func TestForgotPasswordController_Handle(t *testing.T) {
	integration.PrepareDB(t)
	gin.SetMode(gin.TestMode)

	ctrl := NewForgotPasswordController()

	t.Run("happy path - returns 202", func(t *testing.T) {
		integration.PrepareDB(t)
		body, _ := json.Marshal(map[string]interface{}{"email": "user@example.com"})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/auth/forgot-password", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		ctrl.Handle(c)

		assert.Equal(t, http.StatusAccepted, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Contains(t, res["message"], "If an account exists")
	})

	t.Run("main fail path - invalid email returns 400", func(t *testing.T) {
		integration.PrepareDB(t)
		body, _ := json.Marshal(map[string]interface{}{"email": "not-an-email"})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/auth/forgot-password", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		ctrl.Handle(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "INVALID_REQUEST", res["code"])
	})
}
