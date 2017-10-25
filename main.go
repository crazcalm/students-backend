package main

import (
	"fmt"
	_ "students/routers"
	"students/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3" // import your required driver
)

func init(){
	//Registering the Database driver
	orm.RegisterDriver("sqlite3", orm.DRMySQL)
	
	//Registering the Database
	orm.RegisterDataBase("default", "sqlite3", "testing.db")
}

func main() {
	//Creting a connection to the database
	o := orm.NewOrm()
	o.Using("default")

	//Creating a test user to mess with
	user := new(models.User)
	user.Name = "Marcus Willock"

	//Creating a test Student
	s := new(models.Student)
	s.ChineseName = "杜明卫"
	s.Pinyin = "du4ming2wei4"
	s.Sex = "Male"
	s.EnglishName = "Marcus Willock"
	s.StudentID = "007"
	s.Class = "Teacher"
	
	//Force automatic table creation
	// Database alias.
	name := "default"
	
	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true
	
	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
	    fmt.Println(err)
	}

	fmt.Println(o.Insert(user))
	fmt.Println(o.Insert(s))
	
	beego.Run()
}
