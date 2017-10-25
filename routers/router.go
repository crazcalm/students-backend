package routers

import (
	"students/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	beego.Router("/hello", &controllers.LearningController{})

	//Messing with students
	beego.Router("/student", &controllers.StudentController{})
}
