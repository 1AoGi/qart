package routers

import (
	"github.com/astaxie/beego"
	"github.com/tautcony/qart/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/render", &controllers.RenderController{}),
		beego.NSRouter("/render/upload", &controllers.UploadController{}),
		beego.NSRouter("/share", &controllers.ShareController{}, "post:CreateShare"),
	)
	beego.AddNamespace(ns)
}
