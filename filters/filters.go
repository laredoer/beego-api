package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
)

var HasPermission = func(ctx *context.Context) {
	beego.Alert("执行中间件")
}

