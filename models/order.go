package models

import (
	//"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

var OrderTypeMap = map[string]string{"normal": "普通订单", "bottom": "打底订单", "exchange": "换量订单"}

type Order struct {
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
	Cuid        int64  `json:"cuid"`
	Typ         string `json:"type" orm:"column(type)"`
	CampaignId  int64 `json:"campaign_id"`
	OrderType    string `json:"order_type"`
	Source      string `json:"source"`
	ScheduleId  int64  `json:"schedule_id"`
	ScheduleSource string `json:"schedule_source"`
	DepartmentGid int64 `json:"department_gid"`
	DepartmentPid int64 `json:"department_pid"`
	DepartmentId int64 `json:"department_id"`
	DepartmentGname string `json:"department_gname" orm:"-"`
	DepartmentPname string `json:"department_pname" orm:"-"`
	DepartmentName string `json:"department_name" orm:"-"`
	Name string `json:"name"`
	DeletedAt string `json:"-" `
	//Tname       string `json:"-" orm:"column(table_name)"`
}

func (i *Order) TableName() string {
	return "order"
}

func (i *Order) Orm() orm.QuerySeter {
	return orm.NewOrm().QueryTable(i.TableName())
}

func (i *Order) GetList()(list []*Order)  {
	o := orm.NewOrm()
	o.Using("rtb")
	o.QueryTable(i.TableName()).All(&list)

	return
}
