package controllers

import (
	"fmt"
	"log"
	"encoding/json"
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
	}
	
	//This is an example of sending json
	c.Data["json"] = &user
	c.ServeJSON()
}

//Post -- Learning how to use post
func (c *LearningController) Post (){
	c.Data["Website"] = "Now website"
	c.Data["Email"] = "marcuswillock@qq.com"
	c.TplName = "index.tpl"

	fmt.Println("Request body below:")
	fmt.Println(c.Ctx.Input.RequestBody)

	//values in request
	var v map[string]string
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)

	//json map
	name := v["name"]
	if name == "" {
		c.Ctx.WriteString("name is empty\n")
	    return
	}

	err = models.NewUser(name)
	if err != nil {
		log.Fatal(err)
	}
}
