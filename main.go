package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/astaxie/beego/logs"
	"github.com/gin-gonic/gin"
)

var (
	workerNum int
	mode      bool
	host      string
	port      int
)

func init() {
	cpuNum := runtime.NumCPU()

	flag.IntVar(&workerNum, "worker", cpuNum, "运行核数")
	flag.BoolVar(&mode, "debug", false, "是否开启开发模式")
	flag.StringVar(&host, "host", "127.0.0.1", "服务器host")
	flag.IntVar(&port, "port", 8001, "服务器port")

	logs.Async()
}

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(workerNum)

	address := fmt.Sprintf("%s:%d", host, port)
	logs.Info("%s 服务器初始化中... 核数:%d 调试模式:%t", address, workerNum, mode)

	if !mode {
		gin.SetMode(gin.ReleaseMode)
		logs.SetLogger(logs.AdapterMultiFile, `{"filename":"log/tdida.log","separate":["critical", "error", "info"]}`)
	}

	router := gin.Default()

	logs.Info("server成功启动")
	router.Run(address)
}
