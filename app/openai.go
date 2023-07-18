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
	"walnut/scheduler"
	"walnut/util"
)

var GPT4 = "gpt-4"
var GPT35 = "gpt-3.5-turbo"
var GPT3516K = "gpt-3.5-turbo-16k"
var URL = "https://api.openai.com/v1/chat/completions"

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
	var isNewChat = false
	//拼接消息体 redis中先获取是否有缓存
	var messages []model.Message
	messageStr, err := rds.Rds.Get(context.Background(), user).Result()
	if err == redis.Nil {
		// 新的对话
		isNewChat = true
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
	//请求open ai
	modelResp := ChatCompletionsReq(messages, isNewChat)
	//处理返回结果
	finishReason := gjson.Get(modelResp, "choices.0.finish_reason").String()
	if finishReason == "function_call" {
		functionName := gjson.Get(modelResp, "choices.0.message.function_call.name").String()
		if functionName == "train_monitor" {
			var arguments HomeTask
			json.Unmarshal([]byte(gjson.Get(modelResp, "choices.0.message.function_call.arguments").String()), &arguments)
			tag := arguments.FromDate + arguments.TrainNumber
			scheduler.Scheduler.Every(30).Minutes().Tag(tag).Do(arguments.Run)
		} else {
			fmt.Println("function name is not support")
		}
		//调用完成后 返回open ai结果
		funMsg := model.Message{
			Role:    "function",
			Content: "function call successful",
			Name:    functionName,
		}
		messages = append(messages, funMsg)
		modelResp = ChatCompletionsReq(messages, false)
	}
	// 普通消息处理
	newMsg := gjson.Get(modelResp, "choices.0.message").String()
	var message model.Message
	json.Unmarshal([]byte(newMsg), &message)
	messages = append(messages, message)
	messageJsonArray, err := json.Marshal(messages)
	if err != nil {
		fmt.Println("json marshal error:", err)
	}
	rds.Rds.Set(context.Background(), user, messageJsonArray, 1*time.Hour)

	return []byte(modelResp)
}

// ChatCompletionsReq 封装openai chat请求
func ChatCompletionsReq(messages []model.Message, isNewChat bool) string {
	// 获取token
	apiKey, _ := rds.Rds.Get(context.Background(), "api_key").Result()
	headers := map[string]string{
		"Content-Type":  "application/json; charset=utf-8",
		"Authorization": "Bearer " + apiKey}
	// 拼接参数
	var requestChat model.Chat
	if isNewChat {
		requestChat = model.Chat{
			Model:       GPT3516K,
			Messages:    messages,
			Temperature: 0.5,
			Functions: []model.Functions{{
				Name:        "train_monitor",
				Description: "help people monitor train ticket for go home",
				Parameters: model.Parameters{
					Type: "object",
					Properties: map[string]interface{}{
						"trainNumber": map[string]string{
							"type":        "string",
							"description": "train number,e.g. G602, G613",
						},
						"fromDate": map[string]string{
							"type":        "string",
							"description": "date str,e.g. 2021-10-01",
						},
						"fromStation": map[string]interface{}{
							"type":        "string",
							"description": "station code",
							"enum":        []string{"TYV", "BJP"},
						},
						"toStation": map[string]interface{}{
							"type":        "string",
							"description": "station code",
							"enum":        []string{"TYV", "BJP"},
						},
					},
					Required: []string{"trainNumber", "fromDate", "fromStation", "toStation"},
				}}},
		}
	} else {
		requestChat = model.Chat{
			Model:       GPT3516K,
			Messages:    messages,
			Temperature: 0.5,
		}
	}
	// 发送请求
	resp := util.HttpReq("POST", URL, headers, requestChat)
	strResp := string(resp)
	// 打印返回
	fmt.Println("open ai resp:", strResp)
	return strResp
}

// List 列出所有模型
func List(c *gin.Context) {
	url := "https://api.openai.com/v1/models"
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("GET")

	apiKey, _ := rds.Rds.Get(context.Background(), "api_key").Result()

	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error in Do:", err)
	}
	fmt.Println("open ai resp:", string(resp.Body()))

	c.Data(http.StatusOK, "application/json; charset=utf-8", resp.Body())
}
