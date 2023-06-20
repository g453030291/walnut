package app

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

func Fmsg(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "invilid request json"})
	}

	c.JSON(http.StatusOK, sendMsg(string(body)))
}

// 飞书认证challenge方法
func challenge(data string) gin.H {
	return gin.H{
		"challenge": gjson.Get(string(data), "challenge").String(),
	}
}

// 接收,回复飞书消息
func sendMsg(body string) gin.H {
	url := "https://open.feishu.cn/open-apis/im/v1/messages"

	content := gjson.Get(body, "event.message.content").String()
	// 收到的消息
	text := gjson.Get(content, "text").String()
	fmt.Printf("收到的消息:%s\n", text)

	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte("test send msg")))
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return nil
	}

	req.Header.Set("Authorization", "Bearer t-"+gjson.Get(body, "event.message.content").String())
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return nil
	}
	defer resp.Body.Close()

	fmt.Printf("resp: %s\n", resp)

	return gin.H{
		"result": text,
	}
}
