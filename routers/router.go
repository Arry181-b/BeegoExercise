package routers

import (
	"BeegoExercise/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/",&controllers.RegisterController{})
    beego.Router("/", &controllers.MainController{})
}

