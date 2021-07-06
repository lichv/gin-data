package crontabs

import (
	"github.com/robfig/cron"
)

func Setup() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		//log.Println("每秒执行一次")
	})

	c.Run()

	return
}
