package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/liusy182/lastchance/controllers:BankingController"] = append(beego.GlobalControllerRouter["github.com/liusy182/lastchance/controllers:BankingController"],
		beego.ControllerComments{
			Method: "ShowAccounts",
			Router: `/banking`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}