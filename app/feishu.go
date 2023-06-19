package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Fmsg(c *gin.Context) {
	var data map[string]interface{}
	c.BindJSON(&data)

	fmt.Printf("data: %v\n", data)

	c.JSON(http.StatusOK, gin.H{
		"challenge": data["challenge"],
	})
}
