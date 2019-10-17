package main

import (
	"github.com/astaxie/beego"
	_ "makeModels/routers"
)

func main() {
	//beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}
