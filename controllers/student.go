package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	"strings"
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

	//Confirm that an id was sent
	if strings.EqualFold(v["id"], "") == true {
		err = fmt.Errorf("JSON missing field id")
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	_, err = strconv.Atoi(v["id"])
	if err != nil {
		log.Println(err)
		c.Data["json"] = "Make sure the 'id' is a string representation of a number"
		c.ServeJSON()
		return
	}

	//I should validate the input... Later

	//try to updated the students
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
	}
	fmt.Println(v)

	//Confirm that an id was sent
	if strings.EqualFold(v["id"], "") == true {
		err = fmt.Errorf("JSON missing field id")
		return //This is wrong. I need to return the error to the user
	}

	id, err := strconv.Atoi(v["id"])
	if err != nil {
		log.Println(err)
		return
	}

	//Try to delete the student
	err = models.DeleteStudent(id)
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
