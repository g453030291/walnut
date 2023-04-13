package config

import (
	"github.com/gin-gonic/gin"
	"walnut/app"
)

func Router() *gin.Engine {
	router := gin.Default()

	health := router.Group("/health")
	{
		health.GET("/ping", app.Ping)
	}

	return router
}
