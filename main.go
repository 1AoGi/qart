package main

import (
	"github.com/astaxie/beego"
	"os"
	_ "github.com/tautcony/qart/routers"
)

func main() {
	for _, argv := range os.Args {
		if argv == "--prod" {
			beego.BConfig.RunMode = "prod"
		}
	}
	beego.Run()
}
