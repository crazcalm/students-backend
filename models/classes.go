package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/validation"
	"log"
	"strconv"
	"strings"
)

//Class -- testing this out
type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//Valid -- Method for validating the data
func (c *Class) Valid(v *validation.Validation) {
	if strings.EqualFold(c.Name, "") {
		err := v.SetError("Name", "Cannot be an empty string")
		if err != nil {
			log.Println(err)
			return
		}
	}
}

//GetClasses -- returns all the classes that have not been deleted
func GetClasses(conn *sql.DB) (classes []Class, err error) {
	defer conn.Close() // nolint: errcheck

	//Get classes
	rows, err := conn.Query(`SELECT id, name FROM class WHERE deleted = false`)

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

	err = rows.Close()
	if err != nil {
		log.Println(err)
		return
	}
	return
}

//UpdateClassName -- updates the name of the class
func UpdateClassName(conn *sql.DB, id string, name string) (err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Close the database connection
	defer conn.Close() // nolint: errcheck

	//Determine if the class exists
	rows, err := conn.Query(`SELECT name FROM class WHERE id = $1`, idInt)
	if err != nil {
		log.Println(err)
		return
	}

	var originalName string
	for rows.Next() {
		err = rows.Scan(&originalName)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if strings.EqualFold(name, "") {
		err = fmt.Errorf("Class does not exist")
		return
	}

	//Update the name
	_, err = conn.Query(`UPDATE class SET name = $2 WHERE id = $1`, idInt, name)
	if err != nil {
		log.Println(err)
		return
	}

	err = rows.Close()
	if err != nil {
		log.Println(err)
		return
	}

	return
}

//DeleteClass -- flips delete flag for class
func DeleteClass(conn *sql.DB, id string) (err error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	// Close the database
	defer conn.Close() // nolint: errcheck

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

	if strings.EqualFold(name, "") {
		err = fmt.Errorf("Class does not exist")
		return
	}

	//Flip the delete flag
	_, err = conn.Query(`UPDATE class SET deleted = true WHERE id = $1`, idInt)
	if err != nil {
		log.Println(err)
		return
	}

	err = rows.Close()
	if err != nil {
		log.Println(err)
		return
	}

	return
}

//NewClass -- adds a new class to the database
func NewClass(conn *sql.DB, name string) (err error) {
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

	//Close database connection
	defer conn.Close() // nolint: errcheck

	row := conn.QueryRow("INSERT INTO class(name) values($1)", c.Name)
	log.Println(row)

	return
}
