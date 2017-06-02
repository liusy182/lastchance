package routers

import (
	"github.com/liusy182/lastchance/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
