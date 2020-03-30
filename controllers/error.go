package controllers

import (
	"mn_log/trace"

	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (i *ErrorController) Error404() {
	traceObj := (&trace.TraceLog{}).Init()

	msg := "接口不存在"
	code := 404
	data := map[string]interface{}{}
	resp := &Response{
		Code:    code,
		Data:    data,
		Msg:     msg,
		TraceID: traceObj.TraceID,
	}
	traceObj.Uri = i.Ctx.Request.RequestURI
	traceObj.Res.Code = code
	traceObj.Res.Data = data
	traceObj.Res.Msg = msg
	traceObj.Res.TraceID = traceObj.TraceID
	traceObj.IP = i.Ctx.Request.RemoteAddr
	traceObj.UserAgent = i.Ctx.Request.UserAgent()
	traceObj.Method = i.Ctx.Request.Method
	traceObj.Uri = i.Ctx.Request.RequestURI
	traceObj.Req = i.Input()
	traceObj.Add().Close()
	i.Data["json"] = &resp
	i.ServeJSON()
}

func (i *ErrorController) Error501() {
	traceObj := (&trace.TraceLog{}).Init()

	msg := "server error"
	code := 500
	data := map[string]interface{}{}
	resp := &Response{
		Code:    code,
		Data:    data,
		Msg:     msg,
		TraceID: traceObj.TraceID,
	}
	traceObj.Uri = i.Ctx.Request.RequestURI
	traceObj.Res.Code = code
	traceObj.Res.Data = data
	traceObj.Res.Msg = msg
	traceObj.Res.TraceID = traceObj.TraceID
	traceObj.IP = i.Ctx.Request.RemoteAddr
	traceObj.UserAgent = i.Ctx.Request.UserAgent()
	traceObj.Method = i.Ctx.Request.Method
	traceObj.Uri = i.Ctx.Request.RequestURI
	traceObj.Req = i.Input()
	traceObj.Add().Close()
	i.Data["json"] = &resp
	i.ServeJSON()

}
func (i *ErrorController) ErrorDb() {
	traceObj := (&trace.TraceLog{}).Init()
	msg := "db error"
	code := 500
	data := map[string]interface{}{}
	resp := &Response{
		Code:    code,
		Data:    data,
		Msg:     msg,
		TraceID: traceObj.TraceID,
	}
	traceObj.Uri = i.Ctx.Request.RequestURI
	traceObj.Res.Code = code
	traceObj.Res.Data = data
	traceObj.Res.Msg = msg
	traceObj.Res.TraceID = traceObj.TraceID
	traceObj.IP = i.Ctx.Request.RemoteAddr
	traceObj.UserAgent = i.Ctx.Request.UserAgent()
	traceObj.Method = i.Ctx.Request.Method
	traceObj.Uri = i.Ctx.Request.RequestURI
	traceObj.Req = i.Input()
	traceObj.Add().Close()
	i.Data["json"] = &resp
	i.ServeJSON()
}
