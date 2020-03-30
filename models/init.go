package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func init() {
	connect("default")
	connect("user")
	connect("rtb")
	orm.RegisterModel(new(User), new(UserLog), new(UserDataRule), new(UserDataRuleExt),new(Order),new(Department))
}

func connect(aliasName string) {

	if aliasName == "" {
		aliasName = "default"
	}

	ck := ""
	if aliasName != "default" {
		ck += aliasName + "_"
	}

	beego.LoadAppConfig("ini", "../conf/app.conf")

	db_source := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s",
		beego.AppConfig.String(ck+"db_user"), beego.AppConfig.String(ck+"db_pwd"), beego.AppConfig.String(ck+"db_host"),
		beego.AppConfig.String(ck+"db_name"), beego.AppConfig.String(ck+"db_charset"))

	db_type := beego.AppConfig.String(ck + "db_type")
	db_max_open, _ := beego.AppConfig.Int(ck + "db_max_open")
	db_max_idle, _ := beego.AppConfig.Int(ck + "db_max_idle")

	db_debug, derr := beego.AppConfig.Bool(ck + "db_debug")
	if derr != nil {
		orm.Debug = false
	} else {
		orm.Debug = db_debug
	}

	orm.RegisterDriver(db_type, orm.DRMySQL) //注册驱动
	orm.RegisterDataBase(aliasName, db_type, db_source)
	orm.SetMaxIdleConns(aliasName, db_max_idle)
	orm.SetMaxOpenConns(aliasName, db_max_open)
}
