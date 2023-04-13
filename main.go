package main

import (
	"walnut/config"
)

func main() {
	r := config.Router()
	r.Run()
}
