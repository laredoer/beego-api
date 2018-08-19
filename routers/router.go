// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api-app/controllers"

	"github.com/astaxie/beego"
	"api-app/filters"
)

func init() {

	beego.InsertFilter("/v1/*", beego.BeforeRouter, filters.HasPermission)

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/version",
			beego.NSInclude(
				&controllers.DefaultController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
