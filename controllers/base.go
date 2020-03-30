package controllers

import (
	"encoding/json"
	"fmt"
	"mn_log/service"
	"mn_log/trace"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//普通数据结构
type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	TraceID int64       `json:"trace_id"`
}
type ResourceId struct {
	DepartmentGid []int64 `json:"departmentGid"`
	DepartmentPid []int64 `json:"departmentPid"`
	DepartmentId  []int64 `json:"departmentId"`
	BrandId       []int64 `json:"brandId"`
	ProductId     []int64 `json:"productId"`
}
type Empty struct {
}

type BaseController struct {
	beego.Controller
	Page Page
}

//分页数据结构
type Page struct {
	PageNo     int64       `json:"page_no"`
	PageSize   int64       `json:"page_size"`
	TotalPage  int64       `json:"total_page"`
	TotalCount int64       `json:"total_count"`
	FirstPage  int64       `json:"first_page"`
	LastPage   int64       `json:"last_page"`
	List       interface{} `json:"list"`
}

func (i *BaseController) GetResourceId() (res ResourceId) {
	redisObj := (&service.Redis{}).Init()
	defer redisObj.Close()
	key := fmt.Sprintf("userStructure:%d", i.getUid())
	val, err := redisObj.Client.Get(key).Result()
	//fmt.Println(key, val, err)
	if err != nil {
		logs.Error(key, err)
		return
	}
	err = json.Unmarshal([]byte(val), &res)
	if err != nil {
		logs.Error(key, "json.Unmarshal()", err)

		return
	}
	return
}
func (i *Page) Set(count int64, pageNo int64, pageSize int64) {

	c := count

	if pageSize == 0 {
		pageSize, _ = beego.AppConfig.Int64("page::size")
	}

	tp := c / pageSize
	if c%pageSize > 0 {
		tp = c/pageSize + 1
	}

	if pageNo < 1 {
		pageNo = 1
	}

	i.TotalPage = tp
	i.FirstPage = 1
	i.LastPage = tp
	i.PageNo = pageNo
	i.TotalCount = c
	i.PageSize = pageSize
}

func (i *Page) Offset() int64 {
	return (i.PageNo - 1) * i.PageSize
}

func (i *Page) Limit() int64 {
	return i.PageSize
}

//成功
func (i *BaseController) success(data interface{}) {
	traceObj := (&trace.TraceLog{}).Init()
	code := 10000
	msg := beego.AppConfig.String("code::" + strconv.Itoa(code))

	if data == nil {
		data = Empty{}
	}

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

//成功
func (i *BaseController) fail(code int) {
	traceObj := (&trace.TraceLog{}).Init()
	msg := beego.AppConfig.String("code::" + strconv.Itoa(code))

	resp := &Response{
		Code:    code,
		Data:    map[string]interface{}{},
		Msg:     msg,
		TraceID: traceObj.TraceID,
	}
	traceObj.Uri = i.Ctx.Request.RequestURI
	traceObj.Res.Code = code
	traceObj.Res.Data = map[string]interface{}{}
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

//获取资源ID
func (i *BaseController) getResourceId() int64 {
	resource_id := i.Ctx.Input.GetData("inner_resource_id")
	return resource_id.(int64)
}

//获取用户ID
func (i *BaseController) getUid() int64 {
	uid := i.Ctx.Input.GetData("inner_uid")
	return uid.(int64)
}

//获取操作UID
func (i *BaseController) getMuid() int64 {
	muid := i.Ctx.Input.GetData("inner_muid")
	return muid.(int64)
}
