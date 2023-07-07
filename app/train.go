package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"net/http"
)

// Home 回家车票监控
// 北京->太原
func Home(c *gin.Context) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI("https://kyfw.12306.cn/otn/leftTicket/query?leftTicketDTO.train_date=2023-07-14&leftTicketDTO.from_station=BJP&leftTicketDTO.to_station=TYV&purpose_codes=ADULT")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error in Do:", err)
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", resp.Body())
}
