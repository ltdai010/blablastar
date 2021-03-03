// @APIVersion 1.0.0
// @Title web Test API
// @Description web has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"blablastar/controllers/postcontroller"
	"blablastar/controllers/starcontroller"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace("/v1",
		web.NSNamespace("/star",
			web.NSInclude(
				&starcontroller.StarController{},
			),
		),
		web.NSNamespace("/post",
			web.NSInclude(
				&postcontroller.PostController{},
			),
		),
	)
	web.AddNamespace(ns)
}
