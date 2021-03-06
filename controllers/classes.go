package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"strings"
	"students/models"
	"students/constants"
)

//ClassesController -- Controller for the classes model
type ClassesController struct {
	beego.Controller
}

//Get -- Returns a JSON object of all the non-deleted classes
func (c *ClassesController) Get() {
	classes, err := models.GetClasses()
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	fmt.Println(classes)

	classesJSON, err := json.Marshal(classes)
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = string(classesJSON)
	c.ServeJSON()
	return
}

//Post --
func (c *ClassesController) Post() {
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

	//Validate user input
	err = ValidateUserInput(v, []string{"name"})
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//Create the new class
	err = models.NewClass(v["name"])
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = constants.SUCCESS
	c.ServeJSON()
	return
}

//Put --
func (c *ClassesController) Put() {
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

	//Validate user input
	err = ValidateUserInput(v, []string{"id", "name"})
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//Confirm that the id was sent
	if strings.EqualFold(v["id"], "") {
		err = fmt.Errorf("JSON field missing id")
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//try to update a class
	err = models.UpdateClassName(v["id"], v["name"])
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = constants.SUCCESS
	c.ServeJSON()
	return

}

//Delete -- Flips the delete flag or for a class
func (c *ClassesController) Delete() {
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

	//Validate user input
	err = ValidateUserInput(v, []string{"id"})
	if err != nil {
		log.Println(err)
		return
	}

	//try to delete a class
	err = models.DeleteClass(v["id"])
	if err != nil {
		log.Println(err)
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = constants.SUCCESS
	c.ServeJSON()
	return
}
