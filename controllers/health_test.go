package controllers

import (
	"github.com/gin-gonic/gin"
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

	if w.Code != http.StatusOK {
		t.Fail()
	}

	if w.Body.String() != StatusResult {
		t.Fail()
	}
}
