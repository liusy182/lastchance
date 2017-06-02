package controllers

import (
	"github.com/astaxie/beego"
)

type SecurityController struct {
	beego.Controller
}

func (c *SecurityController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
