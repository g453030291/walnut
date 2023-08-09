package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HttpReq 请求封装
func HttpReq(method string, url string, headers map[string]string, param any) []byte {
	client := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var byteParam io.Reader
	if param != nil {
		jsonBytes, err := json.Marshal(param)
		if err != nil {
			fmt.Println("HttpReq json.Marshal error:", err)
		}
		byteParam = bytes.NewBuffer(jsonBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, byteParam)

	if err != nil {
		fmt.Printf("HttpReq http.NewRequest error:%v\n", err)
	}

	//set header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("HttpReq client.Do error:%v\n", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("HttpReq resp.Body.Close error:", err)
		}
	}()

	//读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("HttpReq io.ReadAll error:", err)
	}

	return body
}
