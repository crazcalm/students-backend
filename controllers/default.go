package controllers

import (
	"fmt"
	"log"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3" //database driver
	"github.com/astaxie/beego"
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
	c.Data["json"] = "nothing here"
	c.ServeJSON()
}

//Post -- Learning how to use post
func (c *LearningController) Post (){
	//Trying to see into the request body
	fmt.Println("Request body below:")
	fmt.Println(c.Ctx.Input.RequestBody)

	//values in request
	var v map[string]string
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	//json map
	name := v["name"]
	if name == "" {
		c.Ctx.WriteString("name is empty\n")
	}
}
