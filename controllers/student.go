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

//Get -- Get all the students that have not been deleted
func (c *StudentController) Get() {
	students, err := models.GetStudents()
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	fmt.Println(students)

	ss, err := json.Marshal(students)
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = string(ss)
	c.ServeJSON()
	return
}

//Put -- Update a student
func (c *StudentController) Put() {
	fmt.Println("Request body below:")
	fmt.Println(c.Ctx.Input.RequestBody)

	//values in the request
	var v map[string]string
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	fmt.Println(v)

	//Validation of user input
	err = ValidateUserInput(v, []string{"id", "chinese_name", "pinyin", "student_id", "class_id", "sex_id"})
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//try to updated the student
	err = models.UpdateStudent(v["id"], v["chinese_name"], v["pinyin"], v["english_name"], v["student_id"], v["class_id"], v["sex_id"])
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

//Delete -- Flips the delete flag for a student
func (c *StudentController) Delete() {
	fmt.Println("Request body below:")
	fmt.Println(c.Ctx.Input.RequestBody)

	//values in request
	var v map[string]string
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}
	fmt.Println(v)

	//Validate user input
	err = ValidateUserInput(v, []string{"id"})
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//Try to delete the student
	err = models.DeleteStudent(v["id"])
	if err != nil {
		log.Println(err)
		c.Data["json"] = fmt.Sprintf(err.Error())
		c.ServeJSON()
		return
	}
	c.Data["json"] = "success"
	c.ServeJSON()
	return
}

//Post -- Learning how to use post
func (c *StudentController) Post() {
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
