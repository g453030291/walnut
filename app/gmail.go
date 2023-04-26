package app

import (
	"walnut/constans"
	"walnut/util"
)

func GetProfile() string {
	url := constans.GMAIL_BASE_URL + "/gmail/v1/users/112702485909708385821/profile"
	return util.HttpGet(url)
}
