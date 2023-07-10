package main

import (
	"walnut/config"
	"walnut/rds"
)

func main() {
	rds.Init()
	config.CronConfig()
	r := config.Router()
	r.Run(":8080")
}
