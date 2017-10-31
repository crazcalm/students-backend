package routers

import (
	"github.com/astaxie/beego"
	"students/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/hello", &controllers.LearningController{})

	//Messing with students
	beego.Router("/student", &controllers.StudentController{})
}
