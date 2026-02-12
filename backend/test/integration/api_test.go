package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
	"rodrigoorlandini/vet-shifter/internal/api"
)

func init() {
	_ = utils.LoadEnvironment()
	if os.Getenv("DATABASE_URL") == "" {
		os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/vet_shifter?sslmode=disable")
	}
	db := database.GetPostgresConnection()
	if err := database.RunMigrations(db); err != nil {
		panic(err)
	}
}

func TestAPI_Integration_HappyPath(t *testing.T) {
	if os.Getenv("DATABASE_URL") == "" {
		t.Skip("DATABASE_URL not set, skipping integration test")
	}
	router := api.NewRouter()

	t.Run("POST /companies returns 201", func(t *testing.T) {
		body := map[string]interface{}{
			"cnpj":         "00000000000100",
			"company_name": "Clínica Teste",
			"owner_name":  "Responsável",
			"email":        "clinica@test.com",
			"phone":        "11999990000",
			"password":     "senha123",
			"consent_lgpd": true,
		}
		raw, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/companies", bytes.NewReader(raw))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("POST /shift-vets returns 201", func(t *testing.T) {
		body := map[string]interface{}{
			"email":        "vet@test.com",
			"phone":        "11988880000",
			"password":     "vet123",
			"full_name":   "Dr. Veterinário",
			"cpf":         "12345678901",
			"crmv_number": "54321",
			"crmv_state":  "SP",
			"specialties": []string{"Clínico"},
			"consent_lgpd": true,
		}
		raw, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/shift-vets", bytes.NewReader(raw))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("POST /auth/login returns 200 with token", func(t *testing.T) {
		body := map[string]string{"email": "clinica@test.com", "password": "senha123"}
		raw, _ := json.Marshal(body)
		req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(raw))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var res map[string]interface{}
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.Nil(t, err)
		assert.NotEmpty(t, res["token"])
		assert.Equal(t, "clinic", res["role"])
	})

	t.Run("GET /shifts returns 200 with list", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/shifts", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		var res map[string]interface{}
		err := json.NewDecoder(w.Body).Decode(&res)
		assert.Nil(t, err)
		assert.NotNil(t, res["shifts"])
	})
}
