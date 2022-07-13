package main

import (
	"github.com/astaxie/beego"
	_ "makeControlModel/routers"

)

func main() {
	//beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}
