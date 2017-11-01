package controllers

import (
	"github.com/astaxie/beego"
)

//MainController the default provided by the Bee tool
type MainController struct {
	beego.Controller
}

//Get -- Get method
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
