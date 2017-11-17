package routers

import (
	"github.com/astaxie/beego"
	"students/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//Messing with students
	beego.Router("/student", &controllers.StudentController{})

	//Messing with classes
	beego.Router("class", &controllers.ClassesController{})
}
