package app

import (
	"io"
	"log"
	"net/http"
)

func Whoer() {
	resp, err := http.Get("https://whoer.net")
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println(string(body))
}
