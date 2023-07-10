package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
	"time"
	"walnut/rds"
	"walnut/util"
)

// Home 回家车票监控
// 北京->太原
func Home(c *gin.Context) {
	c.Data(http.StatusOK, "application/json; charset=utf-8", GoHome())
}

func GoHome() []byte {
	fmt.Println("运行时间:" + time.Now().Format("2006-01-02 15:04:05"))
	headers := map[string]string{
		"Cookie": "_jc_save_fromDate=2023-07-14; _jc_save_toDate=2023-07-04",
	}
	result := util.HttpReq("GET", "https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=2023-07-14&leftTicketDTO.from_station=BJP&leftTicketDTO.to_station=TYV&purpose_codes=ADULT", headers, nil)

	respData := string(result)
	httpStatus := gjson.Get(respData, "httpstatus")
	if httpStatus.Int() != 200 {
		print("12306接口请求失败")
	}

	trainResult := gjson.Get(respData, "data.result")

	trainResult.ForEach(func(key, value gjson.Result) bool {
		dataStr := strings.Split(value.String(), "|")
		// 二等 一等 商务
		if dataStr[3] == "G613" {
			alertMsg := map[string]interface{}{
				"msg_type": "text",
				"content": map[string]string{
					"text": fmt.Sprintf("车次:%s 商务:%s 一等:%s 二等:%s", dataStr[3], dataStr[32], dataStr[31], dataStr[30]),
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
