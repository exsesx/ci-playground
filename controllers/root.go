package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RootController struct{}

func (h RootController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
