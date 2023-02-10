package init

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// 日志
func init() {
	// 日志输出到文件
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/app.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":30}`)

	// 关闭控制台输出
	isDev := beego.AppConfig.String("runmode") == "dev"
	if !isDev {
		beego.BeeLogger.DelLogger(logs.AdapterConsole)
	}

	logs.EnableFuncCallDepth(true) // 输出文件名和行号
	logs.SetLogFuncCallDepth(3)    // 调用的层级
	logs.Async()                   // 异步输出日志
}
