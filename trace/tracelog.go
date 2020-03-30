package trace

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/astaxie/beego/logs"
)

type TraceLog struct {
	TraceID   int64
	Uri       string       `json:"uri"`
	Res       ResponseInfo `json:"res"`
	Req       url.Values   `json:"req"`
	IP        string       `json:"ip"`
	UserAgent string       `json:"user_agent"`
	Method    string       `json:"method"`
}

//普通数据结构
type ResponseInfo struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
	TraceID int64       `json:"trace_id"`
}

func (i *TraceLog) Init() *TraceLog {
	i.TraceID = time.Now().UnixNano()
	return i
}
func (i *TraceLog) Add() *TraceLog {
	jsonStu, _ := json.Marshal(i)
	logs.Info(string(jsonStu))
	return i
}
func (i *TraceLog) Close() bool {
	i.TraceID = 0
	return true
}
