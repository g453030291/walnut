package config

import (
	"github.com/gin-gonic/gin"
	"walnut/app"
)

func Router() *gin.Engine {
	//创建gin
	r := gin.New()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "success",
		})
	})

	health := r.Group("/health")
	{
		health.GET("/ping", app.Ping)
	}

	fmsg := r.Group("/fmsg")
	{
		fmsg.POST("", app.Fmsg)
	}

	openai := r.Group("/openai")
	{
		openai.POST("", app.MakingRequest)
		openai.GET("/list", app.List)
	}

	home := r.Group("/home")
	{
		home.POST("", app.Home)
		home.POST("/remove", app.RemoveTask)
	}

	return r
}
