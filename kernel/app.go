package kernel

import (
	"kwd/app/crontab"
	"kwd/kernel/api"
	"kwd/kernel/authorize"
	"kwd/kernel/cmd"
	"kwd/kernel/config"
	"kwd/kernel/database"
	"kwd/kernel/logger"
	"kwd/kernel/snowflake"
)

var services = []func(){
	config.InitConfig,
	api.InitApi,
	logger.InitLogger,
	database.InitDatabase,
	database.InitRedis,
	authorize.InitCasbin,
	snowflake.InitSnowflake,
	crontab.InitCrontab,
	cmd.InitCmd,
}

func Bootstrap() {

	for _, item := range services {
		item()
	}

}
