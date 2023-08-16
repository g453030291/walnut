package util

import (
	"context"
	"fmt"
	"github.com/apistd/uni-go-sdk"
	unisms "github.com/apistd/uni-go-sdk/sms"
	"walnut/model"
	"walnut/rds"
)

func SendSms(sms model.Sms) (res *uni.UniResponse) {
	accessKey, _ := rds.Rds.Get(context.Background(), "unisms_access_key").Result()
	signature, _ := rds.Rds.Get(context.Background(), "unisms_signature").Result()
	client := unisms.NewClient(accessKey)

	message := unisms.BuildMessage()
	message.SetTo(sms.Tel)
	message.SetSignature(signature)
	message.SetTemplateId(sms.TemplateId)
	message.SetTemplateData(sms.Param)

	// 发送短信
	res, err := client.Send(message)
	fmt.Println("number:%s, param:%s", sms.Tel, sms.Param)
	if err != nil {
		fmt.Println("sendSms err:%s", err)
		return
	}
	return res
}
