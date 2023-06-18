package config

import (
	"github.com/gin-gonic/gin"
	"walnut/app"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	health := router.Group("/health")
	{
		health.GET("/ping", app.Ping)
	}

	fmsg := router.Group("/fmsg")
	{
		fmsg.GET("/", app.Fmsg)
	}

	return router
}
