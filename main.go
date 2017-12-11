package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
	"tdida.me/apis"
	"tdida.me/config"
	"tdida.me/model"
)

var mode bool

func init() {
	flag.BoolVar(&mode, "debug", false, "是否开启开发模式")
	logs.Async()
}

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	if !mode {
		gin.SetMode(gin.ReleaseMode)
		// logs.SetLogger(logs.AdapterMultiFile, `{"filename":"log/tdida.log","separate":["critical", "error", "info"]}`)
	}

	config.NewConfig(gin.Mode())

	model.InitModel()

	address := fmt.Sprintf("%s:%s", config.ConfigFile.String("httphost"), config.ConfigFile.String("httpport"))
	logs.Info("%s 服务器初始化中... 核数:%d 调试模式:%t", address, runtime.NumCPU(), mode)

	router := gin.Default()

	v1 := router.Group("/v1")

	{
		v1.GET("/", apis.Index)
	}

	logs.Info("server成功启动")
	router.Run(address)
}
