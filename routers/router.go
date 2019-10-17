package routers

import (
	"github.com/astaxie/beego"
	"makeModels/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/makeModels", &controllers.MakeModelsController{})
	beego.Router("/makeModels/postMakeModels", &controllers.MakeModelsController{}, "post:PostMakeModels")
	beego.Router("/makeControl", &controllers.MakeControlController{})
	beego.Router("/makeControl/postControl", &controllers.MakeControlController{}, "post:PostMakeControl")
}
