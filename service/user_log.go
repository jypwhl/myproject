package service

import (
	"encoding/json"
	"mn_log/models"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/logs"
)

type UserLog struct {
}
type UserLogQueue struct {
	Ip        string `json:"ip"`
	Cuid      int64  `json:"cuid"`
	Type      string `json:"type"`
	PkId      int64  `json:"pk_id"`
	TableName string `json:"table_name"`
	Content   string `json:"content"`
	UserId    int64  `json:"user_id"`
	TraceId   string `json:"trace_id"`
	CreatedAt int64  `json:"created_at"`
}

func (i *UserLog) UserLogThread() {
	redisObj := (&Redis{}).Init()
	defer redisObj.Close()
	queneKey := "userLogQueue"
	for {
		res, err := redisObj.Client.RPop(queneKey).Result()
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}
		logs.Info(queneKey, res)
		p := (UserLogQueue{})
		err = json.Unmarshal([]byte(res), &p)
		if err == nil {
			i.Add(p)
		} else {
			logs.Error(queneKey+" json err:", err)
		}
	}
}
func (i *UserLog) Add(param UserLogQueue) {
	userLog := (&models.UserLog{})
	userLog.Ip = param.Ip
	userLog.Cuid = param.Cuid
	userLog.Typ = param.Type
	userLog.PkId = param.PkId
	userLog.Tname = param.TableName
	userLog.Content = param.Content
	userLog.UserId = param.UserId
	userLog.CreatedAt = param.CreatedAt
	userLog.TraceId = param.TraceId

	id, err := orm.NewOrm().Insert(userLog)
	if err != nil || id <= 0 {
		logs.Error("UserLog:Add:err", err)
	}
	return
}
