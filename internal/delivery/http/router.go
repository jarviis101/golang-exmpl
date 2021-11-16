package http

import (
	"github.com/gin-gonic/gin"
	"prj/pkg/auth"
	"prj/pkg/middlewares"
)

func New() *gin.Engine {
	var router = gin.New()
	router.Use(middlewares.JsonHeaderMiddleware())
	router.POST("/login", auth.Login)

	return router
}
