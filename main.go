package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "github.com/tautcony/qart/routers"
	"os"
)

func main() {
	for _, argv := range os.Args {
		if argv == "--prod" {
			beego.BConfig.RunMode = "prod"
		}
	}
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
