package controllers

import (
	"fmt"
	"log"
	"encoding/json"
	"students/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//StudentController -- What I am using to learn beego
type StudentController struct {
	beego.Controller
}

//Get -- Learning this stuff
func (c *StudentController) Get() {
	//Creting a connection to the database
	o := orm.NewOrm()
	o.Using("default")

	//Searching for a student
	student := models.Student{ID:1}

	err := o.Read(&student)
	
	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	}
	
	//This is an example of sending json
	c.Data["json"] = &student
	c.ServeJSON()
}

//Post -- Learning how to use post
func (c *StudentController) Post (){
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
	name := v["chinese_name"]
	if name == "" {
		c.Ctx.WriteString("name is empty\n")
	    return
	}

	err = models.NewStudent(v["chinese_name"], v["pinyin"], v["english_name"], v["student_id"], v["class"], v["sex"])
	if err != nil {
		log.Fatal(err)
	}
}
