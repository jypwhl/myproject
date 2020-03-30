package controllers

import (
	"mn_log/models"
	"time"

	"github.com/syyongx/php2go"

	"github.com/astaxie/beego/orm"
)

type LogController struct {
	BaseController
}

/**
 * 概览统计
 */
func (i *LogController) List() {
	rid := i.GetResourceId()
	page, _ := i.GetInt64("page", 1)
	page_size, _ := i.GetInt64("page_size", 10)
	department_gid, _ := i.GetInt64("department_gid")
	department_pid, _ := i.GetInt64("department_pid")
	department_id, _ := i.GetInt64("department_id")
	brand_id, _ := i.GetInt64("brand_id")
	product_id, _ := i.GetInt64("product_id")
	cuid, _ := i.GetInt64("cuid")
	typ := i.GetString("type")
	start_date := i.GetString("start_date")
	end_date := i.GetString("end_date")
	var list []*models.UserLog

	o := orm.NewOrm()
	qs := o.QueryTable((&models.UserLog{}).TableName())
	condition := orm.NewCondition()
	// condition = condition.AndCond(orm.NewCondition().And("user_id", i.getUid()))

	if typ != "" {
		checkMap := models.TypeMap
		_, ok := checkMap[typ]
		if ok {
			condition = condition.AndCond(orm.NewCondition().And("type", typ))

		}
	}

	if cuid > 0 {
		condition = condition.AndCond(orm.NewCondition().And("cuid", cuid))
	} else {
		user_data_condition := orm.NewCondition()
		var id_ary []int64
		if department_gid > 0 && len(rid.DepartmentGid) > 0 {
			id_ary = append(id_ary, department_gid)
		}
		if department_pid > 0 && len(rid.DepartmentPid) > 0 {
			id_ary = append(id_ary, department_pid)
		}
		if department_id > 0 && len(rid.DepartmentId) > 0 {
			id_ary = append(id_ary, department_id)
		}
		if brand_id > 0 && len(rid.BrandId) > 0 {
			id_ary = append(id_ary, brand_id)
		}
		if product_id > 0 && len(rid.ProductId) > 0 {
			id_ary = append(id_ary, product_id)
		}
		if len(id_ary) > 0 {
			user_data_condition = user_data_condition.AndCond(orm.NewCondition().And("org_id__in", id_ary))
		} else {
			user_data_condition = user_data_condition.AndCond(orm.NewCondition().And("uid", i.getUid()))
		}
		main := orm.NewOrm()
		main.Using("user")
		var id_list []*models.UserDataRuleExt
		user_data_rule_qs := main.QueryTable((&models.UserDataRuleExt{}).TableName())
		user_data_rule_qs = user_data_rule_qs.SetCond(user_data_condition)
		user_data_rule_qs.GroupBy("uid").All(&id_list)
		if len(id_list) > 0 {
			var ids []int64
			for _, x := range id_list {
				ids = append(ids, x.Uid)
			}
			condition = condition.AndCond(orm.NewCondition().And("cuid__in", ids))

		} else {
			condition = condition.AndCond(orm.NewCondition().And("cuid", 0))

		}

	}

	if start_date != "" {
		if end_date == "" {
			end_date = php2go.Date("2006-01-02", time.Now().Unix())
		}
		start_time, _ := php2go.Strtotime("2006-01-02 15:04:05", start_date+" 00:00:00")
		end_time, _ := php2go.Strtotime("2006-01-02 15:04:05", end_date+" 23:59:59")

		cond_date := orm.NewCondition().And("created_at__gte", start_time).And("created_at__lte", end_time)
		condition = condition.AndCond(cond_date)
	}
	qs = qs.SetCond(condition)
	count, _ := qs.Count()

	i.Page.Set(count, page, page_size)
	qs.OrderBy("-id").Limit(i.Page.PageSize, i.Page.Offset()).All(&list)
	var cuidAry []int64
	if len(list) > 0 {

		for _, x := range list {
			cuidAry = append(cuidAry, x.Cuid)
			x.CreatedTime = php2go.Date("2006-01-02 15:04:05", x.CreatedAt)
		}

		userMap := (&models.User{}).GetList(cuidAry)
		for _, x := range list {
			createdUser, ok := userMap[x.Cuid]
			if ok {
				x.CreatedUser = createdUser
			} else {
				x.CreatedUser = ""

			}
		}
	}

	i.Page.List = list

	i.success(i.Page)
}
