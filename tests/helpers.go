package tests

import (
	"encoding/json"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

func performRequest(server *gin.Engine, method string, path string, body any) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(method, path, strings.NewReader(string(jsonBody)))
	res := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	server.ServeHTTP(res, req)

	return res
}
