package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const StatusResult = "Working!"

func TestHealthController_Status(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	healthController := HealthController{}
	healthController.Status(ctx)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, StatusResult, w.Body.String())
}
