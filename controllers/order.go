package controllers

import (
	"mn_log/models"
	"github.com/astaxie/beego/orm"
)


type OrderController struct {
	BaseController
}

/**
 * 订单列表
 */
func (i *OrderController) List() {
	//rid:=i.GetResourceId()
	order_type := i.GetString("order_type")
	page,_:= i.GetInt64("page")
	page_size,_:= i.GetInt64("page_size")
	schedule_id,_:= i.GetInt64("schedule_id")
	search:=i.GetString("search")
	/*page,_:= i.GetInt64("page")
	page_size,_:= i.GetInt64("page_size")
	department_gid,_:= i.GetInt64("department_gid")
	department_pid,_:= i.GetInt64("department_pid")
	department_id,_:= i.GetInt64("department_id")
	brand_id,_:= i.GetInt64("brand_id")
	product_id,_:= i.GetInt64("product_id")
	schedule_id,_:= i.GetInt64("schedule_id")
	campaign_id,_:= i.GetInt64("campaign_id")
	campaign_id,_:= i.GetInt64("campaign_id")
	order_type := i.GetString("order_type")*/
	var list []*models.Order

	o:=orm.NewOrm()
	o.Using("rtb")
	qs:=o.QueryTable((&models.Order{}).TableName())
	condition:=orm.NewCondition()

    if order_type!=""{
		checkMap := models.OrderTypeMap
    	_,ok:=checkMap[order_type]
    	if ok {
			condition = condition.AndCond(orm.NewCondition().And("order_type", order_type))
		}
	}
	if schedule_id>0 {
		condition = condition.AndCond(orm.NewCondition().And("schedule_id", schedule_id))
	}

	qs = qs.SetCond(condition)
	if search!=""{
		qs=qs.Filter("name__contains", search).Filter("id__gte", 69)
	}
	qs=qs.Filter("deleted_at__isnull", true)

	qs=qs.GroupBy("campaign_id")
	// GROUP BY id,age
	count, _ := qs.Count()
	i.Page.Set(count, page, page_size)
	qs.OrderBy("-id").Limit(i.Page.PageSize, i.Page.Offset()).All(&list)
	i.Page.List=list
	i.success(i.Page)

	//orderList :=(&models.Order{}).GetList()

/*	departmentMap := (&models.Department{}).GetList()
	for _, x := range orderList {
		DepartmentGname, ok := departmentMap[x.DepartmentGid]
		if ok {
			x.DepartmentGname = DepartmentGname
		} else {
			x.DepartmentGname = ""
		}
		DepartmentPname, ok := departmentMap[x.DepartmentPid]
		if ok {
			x.DepartmentPname = DepartmentPname
		} else {
			x.DepartmentPname = ""
		}
		DepartmentName, ok := departmentMap[x.DepartmentId]
		if ok {
			x.DepartmentName = DepartmentName
		} else {
			x.DepartmentName = ""
		}
	}

    i.success(orderList)*/

}
