package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

type Department struct {
	Id       int64  `json:"-"`
	Name    string `json:"name"`
}

func (i *Department) TableName() string {
	return "department"
}

/*func (i *Department) GetList(did []int64) map[int64]string {
	var list []Department
	departmentMap := make(map[int64]string)

	o := orm.NewOrm()

	o.Using("user")

	if len(did)==0{
		_, err := o.QueryTable(i.TableName()).All(&list)
		if err != nil {
			logs.Error("models.department.GetList:", err.Error())
		}
	}else{
		_, err := o.QueryTable(i.TableName()).Filter("id__in", did).All(&list)
		if err != nil {
			logs.Error("models.department.GetList:", err.Error())
		}
	}



	for _, v := range list {
		departmentMap[v.Id] = v.Name
	}

	return departmentMap
}
*/
/*func (i *Department) GetList() map[int64]string {
	var list []Department
	departmentMap := make(map[int64]string)

	o := orm.NewOrm()

	o.Using("user")

	if len(did)==0{
		_, err := o.QueryTable(i.TableName()).All(&list)
		if err != nil {
			logs.Error("models.department.GetList:", err.Error())
		}
	}else{
		_, err := o.QueryTable(i.TableName()).Filter("id__in", did).All(&list)
		if err != nil {
			logs.Error("models.department.GetList:", err.Error())
		}
	}



	for _, v := range list {
		departmentMap[v.Id] = v.Name
	}

	return departmentMap
}
*/
func (i *Department) GetList() map[int64]string {
	var list []Department
	departmentMap := make(map[int64]string)

	o := orm.NewOrm()
	o.Using("user")
	_, err := o.QueryTable(i.TableName()).All(&list)
	if err != nil {
		logs.Error("models.department.GetList:", err.Error())
	}
	for _, v := range list {
		departmentMap[v.Id] = v.Name
	}
	return departmentMap
}


