package models

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"log"
	"strconv"
	"strings"
	"students/db"
)

//Class -- testing this out
type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name", valid:"Required"`
}

//Valid -- Method for validating the data
func (c *Class) Valid(v *validation.Validation) {
	if strings.EqualFold(c.Name, "") {
		v.SetError("Name", "Cannot be an empty string")
	}
}

//GetClasses -- returns all the classes that have not been deleted
func GetClasses() (classes []Class, err error) {
	conn := db.DB()
	defer conn.Close()

	//Get classes
	rows, err := conn.Query(`SELECT id, name FROM class WHERE deleted = false`)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() {
		var class Class
		err = rows.Scan(&class.ID, &class.Name)
		if err != nil {
			log.Println(err)
			return
		}
		classes = append(classes, class)
	}
	return
}

//UpdateClassName -- updates the name of the class
func UpdateClassName(id string, name string) (err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	conn := db.DB()
	defer conn.Close()

	//Determine if the class exists
	rows, err := conn.Query(`SELECT name FROM class WHERE id = $1`, idInt)

	var originalName string
	for rows.Next() {
		err = rows.Scan(&originalName)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if strings.EqualFold(name, "") == true {
		err = fmt.Errorf("Class does not exist")
		return
	}

	//Update the name
	_, err = conn.Query(`UPDATE class SET name = $2 WHERE id = $1`, idInt, name)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

//DeleteClass -- flips delete flag for class
func DeleteClass(id string) (err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	conn := db.DB()
	defer conn.Close()

	//Determine if the class exists
	rows, err := conn.Query(`SELECT name FROM class WHERE id = $1`, idInt)
	if err != nil {
		log.Println(err)
		return
	}

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if strings.EqualFold(name, "") == true {
		err = fmt.Errorf("Class does not exist")
		return
	}

	//Flip the delete flag
	_, err = conn.Query(`UPDATE class SET deleted = true WHERE id = $1`, idInt)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

//NewClass -- adds a new class to the database
func NewClass(name string) (err error) {
	c := new(Class)
	c.Name = name

	//Validate the data coming in
	valid := validation.Validation{}

	//Validate New Class
	b, err := valid.Valid(c)
	if err != nil {
		return
	}

	if !b {
		// validation does not pass
		for _, vError := range valid.Errors {
			log.Println(vError.Key, vError.Message)
			err = fmt.Errorf("%s: %s", vError.Key, vError.Message)
			return
		}
	}

	//Database connection
	conn := db.DB()
	defer conn.Close()

	row := conn.QueryRow("INSERT INTO class(name) values($1)", c.Name)
	log.Println(row)
	return
}
