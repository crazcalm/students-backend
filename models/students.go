package models

import (
	"fmt"
	"log"
	"strings"
	"time"
    _ "github.com/mattn/go-sqlite3" //orm driver
    "github.com/astaxie/beego/validation"
)

//Student struct to hold student information
type Student struct {
	ID			int			`json:"-", orm:"auto", valid:"Required"`
	ChineseName string 		`csv:"chinese_name", json:"chinese_name", valid:"Required"`
	Pinyin      string 		`csv:"pinyin", json:"pinyin", valid:"Required"`
	EnglishName string 		`csv:"english_name", json:english_name, valid:"Required"`
	StudentID   string 		`csv:"student_id", json:"student_id", orm:"student_id", valid:"Required"`
	Class		*Class		`orm:"rel(fk), json:"class", valid:"Required""`
	Sex			string		`json:"-", valid:"Required"`
	Created 	time.Time 	`orm:"auto_now_add;type(datetime)"`
	Updated 	time.Time 	`orm:"auto_now;type(datetime)"`
}

// Valid - If your struct implemented interface `validation.ValidFormer`
// When all tests in StructTag succeed, it will execute Valid function for custom validation
func (s *Student) Valid(v *validation.Validation) {
	//Check for empty strings
	if strings.EqualFold(s.EnglishName, "") == true {
		v.SetError("English Name", "Cannot be empty")
	}
	if strings.EqualFold(s.ChineseName, "") == true {
		v.SetError("Chinese Name", "Cannot be empty")
	}
	if strings.EqualFold(s.Pinyin, "") == true {
		v.SetError("Pinyin", "Cannot be empty")
	}
	if strings.EqualFold(s.StudentID, "") == true {
		v.SetError("Student ID", "Cannot be empty")
	}
	if strings.EqualFold(s.Sex, "") == true {
		v.SetError("Sex", "Cannot be empty")
	}

	// Limit the sexes to male of female
	if strings.Contains(s.Sex, "male") == false && strings.Contains(s.Sex, "female") == false {
		fmt.Println(s.Sex)
		v.SetError("Sex_options", "Can only by 'male' or 'female'")
	}
}

//NewStudent Adds a new student to the database
func NewStudent(cName, pinyin, eName, sID, class, sex string) (err error) {
	//Create Student
	s := new(Student)
	s.ChineseName = cName
	s.Pinyin = pinyin
	s.EnglishName = eName
	s.StudentID = sID
	s.Sex = sex

	//Initialize validation object
	valid := validation.Validation{}

	//Validate New Student
	b, err := valid.Valid(s)
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
	return
}
