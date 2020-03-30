package main

import (
	"mn_log/controllers"
	_ "mn_log/routers"
	"mn_log/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	//beego.SetStaticPath("down","download")
	beego.ErrorController(&controllers.ErrorController{})
	//logs.Reset()
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/service.log","daily":true,"maxdays":30,"separate":["info","error","debug"]}`)
	logs.EnableFuncCallDepth(true)
	beego.BConfig.Log.AccessLogs = false
	go (&service.UserLog{}).UserLogThread()
	beego.Run()
}
