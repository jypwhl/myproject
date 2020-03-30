package models

import (
	"github.com/astaxie/beego/orm"
)

var TypeMap = map[string]string{"marketing-clouds": "营销云", "data-clouds": "数据云", "finance": "财务", "user": "个人中心"}

type UserLog struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	Cuid        int64  `json:"cuid"`
	Typ         string `json:"type" orm:"column(type)"`
	Ip          string `json:"ip"`
	Content     string `json:"content"`
	CreatedAt   int64  `json:"-"`
	TraceId     string `json:"-"`
	Tname       string `json:"-" orm:"column(table_name)"`
	PkId        int64  `json:"-"`
	CreatedTime string `json:"created_time" orm:"-"`
	CreatedUser string `json:"created_user" orm:"-"`
}

func (i *UserLog) TableName() string {
	return "user_log"
}

func (i *UserLog) Orm() orm.QuerySeter {
	return orm.NewOrm().QueryTable(i.TableName())
}

func (i *UserLog) GetList(ids []string) (list []*UserLog) {
	i.Orm().Filter("id__in", ids).All(&list)
	return
}
