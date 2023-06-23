package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
	"walnut/model"
	"walnut/rds"
)

func Fmsg(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invilid request json"})
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", []byte(sendMsg(string(body))))
}

// 飞书认证challenge方法
func challenge(data string) gin.H {
	return gin.H{
		"challenge": gjson.Get(string(data), "challenge").String(),
	}
}

// 接收,回复飞书消息
func sendMsg(body string) string {
	url := "https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=open_id"

	content := gjson.Get(body, "event.message.content").String()
	// 收到的消息
	text := gjson.Get(content, "text").String()
	fmt.Printf("收到的消息:%s\n", text)

	// 创建一个HTTP请求
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")

	req.Header.Set("Authorization", "Bearer "+tenantToken())
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	//发送消息
	chatResp := Chat(text)
	toText := gjson.Get(string(chatResp), "choices.0.message.content").String()
	m := model.Content{
		Text: toText,
	}

	jsonText, _ := json.Marshal(m)

	requestData := model.SendMsg{
		ReceiveId: gjson.Get(body, "event.sender.sender_id.open_id").String(),
		MsgType:   "text",
		Content:   string(jsonText),
	}
	jsonData, _ := json.Marshal(requestData)
	fmt.Printf("发送的消息:%s\n", string(jsonData))

	req.SetBodyString(string(jsonData))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error in Do:", err)
	}

	fmt.Println(resp.StatusCode())
	fmt.Println(string(resp.Body()))

	return string(resp.Body())
}

func tenantToken() string {
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	appId, _ := rds.Rds.Get(context.Background(), "app_id").Result()
	appSecret, _ := rds.Rds.Get(context.Background(), "app_secret").Result()

	req.SetBodyString(`{"app_id": "` + appId + `","app_secret": "` + appSecret + `"}`)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error in Do:", err)
	}

	return gjson.Get(string(resp.Body()), "tenant_access_token").String()
}
