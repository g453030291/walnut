package config

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"walnut/app"
)

func CronConfig(scheduler *gocron.Scheduler) {
	// Every starts the job immediately and then runs at the
	// specified interval
	job, err := scheduler.Every(10).Minutes().Do(app.GoHome())
	fmt.Println(job)
	if err != nil {
		fmt.Println(err)
	}
}
