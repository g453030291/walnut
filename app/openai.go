package app

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"net/http"
	"walnut/rds"
)

var GPT4 = "gpt-4"
var GPT35 = "gpt-3.5-turbo"

func MakingRequest(c *gin.Context) {
	url := "https://api.openai.com/v1/chat/completions"
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.SetMethod("POST")

	apiKey, _ := rds.Rds.Get(context.Background(), "api_key").Result()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	req.SetBodyString(`{"model": "gpt-3.5-turbo",
							"messages": [{"role": "user", "content": "Say this is a test!"}],
							"temperature":0.7}`)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("Error in Do:", err)
	}

	//print(string(resp.Body()))

	c.Data(http.StatusOK, "application/json; charset=utf-8", resp.Body())
}
