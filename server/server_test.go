package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const HealthCheckResult = "Working!"

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	w := performRequest(router, "GET", "/health")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, HealthCheckResult, w.Body.String())
}

func TestNewRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()

	assert.NotNil(t, router)
}
