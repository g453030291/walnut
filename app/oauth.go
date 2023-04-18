package app

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Oauth() {
	resp, err := http.Get("https://www.baidu.com")

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(body))
}
