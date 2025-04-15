package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yangliang0514/go-rest-api/models"
)

func TestSignup(t *testing.T) {

	user := models.User{
		Name:     "Kevin Yang",
		Email:    "kevin@example.com",
		Password: "supersecretpassword",
	}

	userJson, _ := json.Marshal(user)

	req := httptest.NewRequest("POST", "/signup", strings.NewReader(string(userJson)))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	server.ServeHTTP(res, req)

	var resBody map[string]string
	err := json.Unmarshal(res.Body.Bytes(), &resBody)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, res.Code)
	assert.Equal(t, "User created successfully", resBody["message"])
}
