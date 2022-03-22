package Untils

import "github.com/robfig/cron"

//添加定时任务
var Cr *cron.Cron

func init() {
	Cr = cron.New()
	Cr.Start()
}
