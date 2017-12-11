package model

import (
	"github.com/astaxie/beego/logs"
	"tdida.me/config"
)

func InitModel() {
	logs.Info(config.ConfigFile.String("site.name"))
	logs.Info(config.ConfigFile.String("appname"))
}
