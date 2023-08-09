package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
	"walnut/constans"
	"walnut/model"
	"walnut/util"
)

// VerificationCode 验证码接口
func VerificationCode(c *gin.Context) {
	var sms model.Sms
	err := c.BindJSON(&sms)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "参数错误",
		})
		return
	}
	rand.Seed(time.Now().Unix())
	sms.TemplateId = "pub_verif_login_ttl"
	sms.Param = map[string]string{
		"code": string(rand.Intn(900000) + 100000),
		"ttl":  "3",
	}
	resp := util.SendSms(sms)
	jsonByte, err := json.Marshal(resp)
	c.Data(http.StatusOK, constans.APPLICATION_JSON, jsonByte)
}
