package models

import (
	"fmt"
	"log"
	"strings"
    "github.com/astaxie/beego/orm"
    "github.com/astaxie/beego/validation"
)

//Student struct to hold student information
type Student struct {
	ID			int		`json:"-", orm:"auto", valid:"Required"`
	ChineseName string 	`csv:"chinese_name", json:"chinese_name", valid:"Required"`
	Pinyin      string 	`csv:"pinyin", json:"pinyin", valid:"Required"`
	EnglishName string 	`csv:"english_name", json:english_name, valid:"Required"`
	StudentID   string 	`csv:"student_id", json:"student_id", orm:"student_id", valid:"Required"`
	Class		string	`json:"class", valid:"Required"`
	Sex			string	`json:"-", valid:"Required"`
}

// Valid - If your struct implemented interface `validation.ValidFormer`
// When all tests in StructTag succeed, it will execute Valid function for custom validation
func (s *Student) Valid(v *validation.Validation) {
	if strings.Compare(s.EnglishName, "") == 0 {
		v.SetError("English Name", "Cannot be empty")
	}
}

//NewStudent Adds a new student to the database
func NewStudent(cName, pinyin, eName, sID, class, sex string) error {
	//Get DB conn
	o := orm.NewOrm()

	//Create Student
	s := new(Student)
	s.ChineseName = cName
	s.Pinyin = pinyin
	s.EnglishName = eName
	s.StudentID = sID
	s.Class = class
	s.Sex = sex

	//Initialize validation object
	valid := validation.Validation{}

	//Validate New Student
	b, err := valid.Valid(s)
	if err != nil {
		return err
	}

	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
		return fmt.Errorf("validation errors were found")
	}

	//Insert into the database
	_, err = o.Insert(s)
	return err
}

func init() {
	// register model
	orm.RegisterModel(new(Student))
}
