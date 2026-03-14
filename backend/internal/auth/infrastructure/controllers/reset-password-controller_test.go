package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	"rodrigoorlandini/vet-shifter/test/integration"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestResetPasswordController_Handle(t *testing.T) {
	t.Parallel()
	integration.PrepareDB(t)
	gin.SetMode(gin.TestMode)

	ctrl := NewResetPasswordController()

	insertResetToken := func(token string, email string, expiresAt time.Time) {
		t.Helper()
		conn := database.GetPostgresConnection()
		q := database.NewQueries(conn)
		ctx := context.Background()
		_, err := q.CreatePasswordResetToken(ctx, queries.CreatePasswordResetTokenParams{
			ID:        uuid.Must(uuid.NewV7()),
			Token:     token,
			Email:     email,
			UserType:  queries.UserTypeCompanyOwner,
			ExpiresAt: expiresAt,
		})
		require.NoError(t, err)
	}

	t.Run("happy path - returns 200", func(t *testing.T) {
		integration.PrepareDB(t)
		token := "valid-reset-token-123"
		insertResetToken(token, "owner@example.com", time.Now().Add(1*time.Hour))
		body, _ := json.Marshal(map[string]interface{}{"token": token, "new_password": "newpassword123"})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/auth/reset-password", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		ctrl.Handle(c)

		assert.Equal(t, http.StatusOK, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "Password updated successfully", res["message"])
	})

	t.Run("main fail path - invalid token returns 400", func(t *testing.T) {
		integration.PrepareDB(t)
		body, _ := json.Marshal(map[string]interface{}{"token": "invalid-or-expired", "new_password": "newpassword123"})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest(http.MethodPost, "/auth/reset-password", bytes.NewReader(body))
		c.Request.Header.Set("Content-Type", "application/json")

		ctrl.Handle(c)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var res map[string]interface{}
		require.NoError(t, json.Unmarshal(w.Body.Bytes(), &res))
		assert.Equal(t, "INVALID_RESET_TOKEN", res["code"])
	})
}
