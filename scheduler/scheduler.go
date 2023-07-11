package scheduler

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"time"
)

var Scheduler *gocron.Scheduler

// CronConfig
// 北京丰台 = FTP
// 太原南   = TNV
// 阳泉北   = YPP
func CronConfig() {
	Scheduler = gocron.NewScheduler(time.UTC)
	Scheduler.TagsUnique()
	Scheduler.StartAsync()
	fmt.Println("go cron init success")
	//_, err := scheduler.Every(10).Minutes().Do(app.GoHome("G613", "2023-07-14", "BJP", "TYV"))
	//if err != nil {
	//	fmt.Println("cron error:", err)
	//}
}
