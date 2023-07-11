package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// HttpReq 请求封装
func HttpReq(method string, url string, headers map[string]string, param map[string]interface{}) []byte {
	client := &http.Client{}
	var byteParam io.Reader
	if param != nil {
		jsonBytes, err := json.Marshal(param)
		if err != nil {
			panic(err)
		}
		byteParam = bytes.NewBuffer(jsonBytes)
	}

	req, err := http.NewRequest(method, url, byteParam)

	if err != nil {
		fmt.Printf("http.NewRequest error:%v\n", err)
	}

	//set header
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	//发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("client.Do error:%v\n", err)
	}
	defer resp.Body.Close()

	//读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}
