package server

import (
	"ci-playground/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	rootController := new(controllers.RootController)

	router.GET("/", rootController.Status)

	return router
}
