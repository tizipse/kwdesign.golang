package crontab

import (
	"github.com/robfig/cron/v3"
)

var crontab *cron.Cron

func InitCrontab() {

	crontab = cron.New()

	register()

	go crontab.Start()

}

func register() {

	//dormitory.CrontabDayPeople(crontab)

}
