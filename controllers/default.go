package controllers

import (
	"fmt"
	"students/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//MainController the default provided by the Bee tool
type MainController struct {
	beego.Controller
}

//LearningController -- What I am using to learn beego
type LearningController struct {
	beego.Controller
}

//Get -- Get method
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//Get -- Learning this stuff
func (c *LearningController) Get() {
	//Creting a connection to the database
	o := orm.NewOrm()
	o.Using("default")

	//Searching for a user
	user := models.User{ID:1}

	err := o.Read(&user)
	
	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user.ID, user.Name)
	}
	
	//This is an example of sending json
	c.Data["json"] = &user
	c.ServeJSON()
}