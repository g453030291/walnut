package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
	"time"
	"walnut/config"
	"walnut/rds"
	"walnut/util"
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
	config.Scheduler.Every(10).Minutes().Do(
		GoHome(gjson.Get(param, "trainNumber").String(),
			gjson.Get(param, "fromDate").String(),
			gjson.Get(param, "fromStation").String(),
			gjson.Get(param, "toStation").String()))
	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("add scheduler success!"))

}

// GoHome 回家车票监控
// fromDate 监控时间
// fromStation 出发站
// toStation 到达站
func GoHome(trainNumber string, fromDate string, fromStation string, toStation string) []byte {
	fmt.Println("运行时间:" + time.Now().Format("2006-01-02 15:04:05"))
	headers := map[string]string{
		"Cookie": "_jc_save_fromDate=" + fromDate + "; _jc_save_toDate=" + fromDate,
	}
	result := util.HttpReq("GET", "https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date="+fromDate+"&leftTicketDTO.from_station="+fromStation+"&leftTicketDTO.to_station="+toStation+"&purpose_codes=ADULT", headers, nil)

	respData := string(result)
	httpStatus := gjson.Get(respData, "httpstatus")
	if httpStatus.Int() != 200 {
		fmt.Println("12306接口请求失败")
	}

	trainResult := gjson.Get(respData, "data.result")

	trainResult.ForEach(func(key, value gjson.Result) bool {
		dataStr := strings.Split(value.String(), "|")
		// 二等 一等 商务
		if dataStr[3] == trainNumber {
			alertMsg := map[string]interface{}{
				"msg_type": "text",
				"content": map[string]string{
					"text": fmt.Sprintf("时间:%s\n车次:%s\n商务:%s\n一等:%s\n二等:%s", fromDate, dataStr[3], dataStr[32], dataStr[31], dataStr[30]),
				},
			}
			url, _ := rds.Rds.Get(context.Background(), "train_alert").Result()
			util.HttpReq("POST", url, nil, alertMsg)
			return false
		}
		return true
	})
	return result
}
