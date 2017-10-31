package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"students/models"
)

//StudentController -- What I am using to learn beego
type StudentController struct {
	beego.Controller
}

//Post -- Learning how to use post
func (c *StudentController) Post() {
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

	err = models.NewStudent(v["chinese_name"], v["pinyin"], v["english_name"], v["student_id"], v["class_id"], v["sex_id"])
	if err != nil {
		log.Println(err)
		c.Data["json"] = fmt.Sprintf(err.Error())
		c.ServeJSON()
		return
	}
	c.Data["json"] = "sucess"
	c.ServeJSON()
	return
}
