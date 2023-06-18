package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Fmsg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
