package main

import (
	"walnut/config"
	"walnut/rds"
	"walnut/scheduler"
)

func main() {
	rds.Init()
	scheduler.CronConfig()
	r := config.Router()
	r.Run(":8080")
}
