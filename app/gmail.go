package app

import (
	"io"
	"log"
	"net/http"
	"walnut/constans"
)

func GetProfile() string {
	resp, err := http.Get(constans.GMAIL_BASE_URL + "/gmail/v1/users/112702485909708385821/profile")

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
