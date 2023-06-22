package main

import (
	"walnut/config"
	"walnut/rds"
)

func main() {
	r := config.Router()
	rds.Init()
	r.Run(":8080")
}
