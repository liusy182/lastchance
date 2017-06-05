package main

import (
	"github.com/astaxie/beego"
	"github.com/liusy182/lastchance/filters"
	_ "github.com/liusy182/lastchance/routers"
)

func init() {
	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("default", "mysql", "root:root@orm_test")
}

func main() {
	beego.InsertFilter("/api/lifecycle", beego.BeforeRouter, filters.BeforeRouterFilter)
	beego.InsertFilter("/api/lifecycle", beego.BeforeExec, filters.BeforeExecFilter)
	beego.InsertFilter("/api/lifecycle", beego.AfterExec, filters.AfterExecFilter)
	beego.InsertFilter("/api/lifecycle", beego.FinishRouter, filters.FinishRouterFilter)

	// o := orm.NewOrm()
	// o.Using("default")

	beego.Run()
}
