package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/tidwall/gjson"
	"github.com/valyala/fasthttp"
	"io"
	"net/http"
	"time"
	"walnut/model"
	"walnut/rds"
)

var GPT4 = "gpt-4"
var GPT35 = "gpt-3.5-turbo"

func MakingRequest(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invilid request json"})
	}
	msg := gjson.GetBytes(body, "msg")
	resp := Chat(msg.String(), uuid.New().String())
	c.Data(http.StatusOK, "application/json; charset=utf-8", resp)
}

func Chat(msg string, user string) []byte {
	//拼接消息体 redis中先获取是否有缓存
	var messages []model.Message
	messageStr, err := rds.Rds.Get(context.Background(), user).Result()
	if err == redis.Nil {
		fmt.Printf("this user:%s message cache is nil:\n", user)
		sysMsg := model.Message{
			Role:    "system",
			Content: "You are a helpful assistant.",
		}
		messages = append(messages, sysMsg)
	} else {
		err := json.Unmarshal([]byte(messageStr), &messages)
		if err != nil {
			fmt.Println("json unmarshal error:", err)
		}
	}
	//追加用户消息
	messages = append(messages, model.Message{
		Role:    "user",
		Content: msg,
	})
	requestChat := model.Chat{
		Model:       "gpt-3.5-turbo",
		Messages:    messages,
		Temperature: 1,
	}
	chatJson, _ := json.Marshal(requestChat)

	url := "https://api.openai.com/v1/chat/completions"
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")

	apiKey, _ := rds.Rds.Get(context.Background(), "api_key").Result()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	req.SetBodyString(string(chatJson))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error in Do:", err)
	}

	newMsg := gjson.Get(string(resp.Body()), "choices.0.message").String()
	var message model.Message
	json.Unmarshal([]byte(newMsg), &message)
	messages = append(messages, message)
	rds.Rds.Set(context.Background(), user, messages, 1*time.Hour)

	return resp.Body()
}
