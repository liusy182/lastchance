package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/liusy182/lastchance/models"
)

type AccountController struct {
	beego.Controller
}

func (c *AccountController) Login() {
	if c.Ctx.Input.IsPost() {
		var loginForm models.LoginForm
		c.ParseForm(&loginForm)
		fmt.Println(loginForm)
	}
	c.TplName = "login.tpl"
}
