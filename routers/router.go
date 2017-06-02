package routers

import (
	"github.com/astaxie/beego"
	"github.com/liusy182/lastchance/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.SecurityController{})
}
