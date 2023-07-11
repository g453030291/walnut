package app

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
	"time"
	"walnut/rds"
	"walnut/util"
)

type HomeTask struct {
	TrainNumber string `json:"trainNumber"`
	FromDate    string `json:"fromDate"`
	FromStation string `json:"fromStation"`
	ToStation   string `json:"toStation"`
}

// Run
// fromDate 监控时间
// fromStation 出发站
// toStation 到达站
func (t *HomeTask) Run() {
	fmt.Println("运行时间:" + time.Now().Format("2006-01-02 15:04:05"))
	headers := map[string]string{
		"Cookie": "_jc_save_fromDate=" + t.FromDate + "; _jc_save_toDate=" + t.FromDate,
	}
	result := util.HttpReq("GET", "https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date="+t.FromDate+"&leftTicketDTO.from_station="+t.FromStation+"&leftTicketDTO.to_station="+t.ToStation+"&purpose_codes=ADULT", headers, nil)

	respData := string(result)
	httpStatus := gjson.Get(respData, "httpstatus")
	if httpStatus.Int() != 200 {
		fmt.Println("12306接口请求失败")
	}

	trainResult := gjson.Get(respData, "data.result")

	trainResult.ForEach(func(key, value gjson.Result) bool {
		dataStr := strings.Split(value.String(), "|")
		// 二等 一等 商务
		if dataStr[3] == t.TrainNumber {
			alertMsg := map[string]interface{}{
				"msg_type": "text",
				"content": map[string]string{
					"text": fmt.Sprintf("时间:%s\n站点:%s-%s\n车次:%s\n商务:%s\n一等:%s\n二等:%s", t.FromDate, t.FromStation, t.ToStation, dataStr[3], dataStr[32], dataStr[31], dataStr[30]),
				},
			}
			url, _ := rds.Rds.Get(context.Background(), "test_alert").Result()
			util.HttpReq("POST", url, nil, alertMsg)
			return false
		}
		return true
	})
}
