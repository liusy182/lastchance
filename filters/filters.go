package filters

import (
	"fmt"

	"github.com/astaxie/beego/context"
)

func BeforeRouterFilter(ctx *context.Context) {
	fmt.Println("BeforeRouterFilter")
}

func BeforeExecFilter(ctx *context.Context) {
	fmt.Println("BeforeExecFilter")
}

func AfterExecFilter(ctx *context.Context) {
	// more for error logging, wont'b e invoked when Controller.Render renders
	// successfully
	fmt.Println("AfterExecFilter")
}

func FinishRouterFilter(ctx *context.Context) {
	// more for error logging, wont'b e invoked when Controller.Render renders
	// successfully
	fmt.Println("FinishRouterFilter")
}
