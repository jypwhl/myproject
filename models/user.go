package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int64  `json:"-"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	JobNum   string `json:"job_num"`
}

func (i *User) TableName() string {
	return "users"
}

func (i *User) GetList(uids []int64) map[int64]string {
	var list []User
	userMap := make(map[int64]string)
	if len(uids) == 0 {
		return userMap
	}

	o := orm.NewOrm()

	o.Using("user")

	_, err := o.QueryTable(i.TableName()).Filter("id__in", uids).All(&list, "id", "nickname", "job_num")
	if err != nil {
		logs.Error("models.User.GetList:", err.Error())
	}
	for _, v := range list {
		if v.Nickname != "" {
			userMap[v.Id] = v.Nickname
		} else {
			userMap[v.Id] = v.JobNum

		}
	}

	return userMap
}
