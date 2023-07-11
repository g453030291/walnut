package app

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"walnut/scheduler"
)

// Home 回家车票监控
// 北京->太原
func Home(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invilid request json"})
	}
	param := string(body)
	//添加定时任务
	tag := gjson.Get(param, "fromDate").String() + gjson.Get(param, "fromDate").String()
	task := &HomeTask{
		TrainNumber: gjson.Get(param, "trainNumber").String(),
		FromDate:    gjson.Get(param, "fromDate").String(),
		FromStation: gjson.Get(param, "fromStation").String(),
		ToStation:   gjson.Get(param, "toStation").String()}
	scheduler.Scheduler.Every(30).Minutes().Tag(tag).Do(task.Run)
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("add scheduler success!"))
}

func RemoveTask(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invilid request json"})
	}
	param := string(body)
	tag := gjson.Get(param, "fromDate").String() + gjson.Get(param, "fromDate").String()
	sError := scheduler.Scheduler.RemoveByTag(tag)
	if sError != nil {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("remove scheduler fail!"))
	} else {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("remove scheduler success!"))
	}
}
