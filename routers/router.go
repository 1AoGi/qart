// @APIVersion 1.0.0
// @Title QArt API
// @Description TO BE FILLED
package routers

import (
	"github.com/astaxie/beego"
	"qart/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/render", &controllers.RenderController{}),
		beego.NSRouter("/render/upload", &controllers.UploadController{}),
	)
	beego.AddNamespace(ns)
}
