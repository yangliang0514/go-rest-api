package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yangliang0514/go-rest-api/models"
)

func TestSignup(t *testing.T) {
	t.Run("create user", func(t *testing.T) {
		res := performRequest(server, "POST", "/signup", "", newUser())

		var resBody map[string]string
		err := json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, "User created successfully", resBody["message"])
	})

	t.Run("create user with existing email", func(t *testing.T) {
		res := performRequest(server, "POST", "/signup", "", newUser())
		assert.Equal(t, http.StatusConflict, res.Code)
	})
}

func TestLogin(t *testing.T) {
	t.Run("login user", func(t *testing.T) {
		user := newUser()
		res := performRequest(server, "POST", "/login", "", user)

		var resBody map[string]any
		err := json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, "Login successful", resBody["message"])
		assert.Equal(t, user.Email, resBody["user"].(map[string]any)["email"])
		assert.Equal(t, user.Name, resBody["user"].(map[string]any)["name"])
		assert.NotNil(t, resBody["token"])
	})

	t.Run("login with incorrect password", func(t *testing.T) {
		user := newUser()
		user.Password = "wrong password"

		res := performRequest(server, "POST", "/login", "", user)

		var resBody map[string]any
		err := json.Unmarshal(res.Body.Bytes(), &resBody)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.Equal(t, "Invalid credentials", resBody["error"])
	})
}

func TestProtectedRoute(t *testing.T) {
	user := newUser()
	userInfo := loginUser(server, user.Email, user.Password)

	t.Run("get events with valid token", func(t *testing.T) {
		res := performRequest(server, "GET", "/events", userInfo.Token, nil)
		assert.Equal(t, http.StatusOK, res.Code)
	})

	t.Run("get events with invalid token", func(t *testing.T) {
		res := performRequest(server, "GET", "/events", "invalid token", nil)
		assert.Equal(t, http.StatusUnauthorized, res.Code)
	})
}

func newUser() *models.User {
	return &models.User{
		Name:     "Kevin Yang",
		Email:    "kevin@example.com",
		Password: "supersecretpassword",
	}
}
