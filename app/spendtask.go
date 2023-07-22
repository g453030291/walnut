package app

import (
	"context"
	"walnut/constans"
	"walnut/model"
	"walnut/rds"
	"walnut/util"
)

func Spend() {
	//组织一段长文本
	//发送给open ai
	message := model.Message{
		Role:    "user",
		Content: constans.GptBestPractices,
	}
	//获取总结结果
	resp := ChatCompletionsReq([]model.Message{message}, false)
	//发送飞书
	testAlert, _ := rds.Rds.Get(context.Background(), "test_alert").Result()
	util.HttpReq("POST",
		testAlert,
		map[string]string{constans.CONTENT_TYPE: constans.APPLICATION_JSON},
		map[string]interface{}{"msg_type": "text", "content": map[string]string{"text": resp}})
}
