package util

import (
	"io"
	"log"
	"net/http"
)

func HttpGet(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	return string(body)
}
