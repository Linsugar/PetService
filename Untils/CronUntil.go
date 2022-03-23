package Untils

import (
	"PetService/ControllerViews"
	"github.com/robfig/cron"
)

//添加定时任务
var Cr *cron.Cron

func init() {
	Cr = cron.New()
	Cr.AddFunc("0 13 17 * * ?", Views.GetArticle)
}
