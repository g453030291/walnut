package config

import (
	"github.com/gin-gonic/gin"
	"walnut/app"
)

func Router() *gin.Engine {
	//创建gin
	r := gin.New()

	//静态资源
	//r.Static("/web", "./web/dist")

	//r.GET("/", func(context *gin.Context) {
	//	context.File("./web/dist/index.html")
	//})

	health := r.Group("/health")
	{
		health.GET("/ping", app.Ping)
	}
	// 飞书
	fmsg := r.Group("/fmsg")
	{
		fmsg.POST("", app.Fmsg)
	}
	// openai
	openai := r.Group("/openai")
	{
		openai.POST("", app.MakingRequest)
		openai.GET("/list", app.List)
		openai.GET("/autoSpend", app.AutoSpend)
		openai.GET("/autoSpendTask", app.AutoSpendTask)
	}
	// train
	home := r.Group("/home")
	{
		home.POST("", app.Home)
		home.POST("/remove", app.RemoveTask)
	}

	return r
}
