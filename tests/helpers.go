package tests

import (
	"encoding/json"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func performRequest(server *gin.Engine, method string, path string, token string, body any) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(method, path, strings.NewReader(string(jsonBody)))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	server.ServeHTTP(res, req)

	return res
}

type UserInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func loginUser(server *gin.Engine, email string, password string) *UserInfo {
	body := map[string]string{
		"email":    email,
		"password": password,
	}

	res := performRequest(server, "POST", "/login", "", body)

	var userInfo UserInfo
	json.Unmarshal(res.Body.Bytes(), &userInfo)

	return &userInfo
}
