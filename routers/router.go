package routers

import (
	"github.com/astaxie/beego"
	"makeControlModel/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/makeModels", &controllers.MakeModelsController{})
	beego.Router("/makeModels/postMakeModels", &controllers.MakeModelsController{}, "post:PostMakeModels")
	beego.Router("/makePretreatmentModels", &controllers.MakePretreatmentModelsController{},"get:GetPretreatment")
	beego.Router("/makePretreatmentModels/postMakePretreatmentModels", &controllers.MakePretreatmentModelsController{}, "post:PostMakePretreatmentModels")
	beego.Router("/makeControl", &controllers.MakeControlController{})
	beego.Router("/makeControl/postControl", &controllers.MakeControlController{}, "post:PostMakeControl")
	beego.Router("/makeColaModels", &controllers.MakeColaModelsController{},"get:GetColaPretreatment")
	beego.Router("/makeColaModels/postMakeColaModels", &controllers.MakeColaModelsController{}, "post:PostMakeColaModels")
}
