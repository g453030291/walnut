package main

import (
	"github.com/go-co-op/gocron"
	"time"
	"walnut/config"
	"walnut/rds"
)

func main() {
	r := config.Router()
	rds.Init()
	scheduler := gocron.NewScheduler(time.UTC)
	config.CronConfig(scheduler)
	r.Run(":8080")
}
