package models

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"log"
	"strconv"
	"strings"
	"students/db"
	"time"
)

//Student struct to hold student information
type Student struct {
	ID          int       `json:"id", valid:"Required"`
	ChineseName string    `csv:"chinese_name", json:"chinese_name", valid:"Required"`
	Pinyin      string    `csv:"pinyin", json:"pinyin", valid:"Required"`
	EnglishName string    `csv:"english_name", json:english_name, valid:"Required"`
	StudentID   string    `csv:"student_id", json:"student_id", orm:"student_id", valid:"Required"`
	ClassID     int       `json:"class_id", valid:"Required"`
	SexID       int       `json:"sex_id", valid:"Required"`
	Created     time.Time `json:"-"`
	Updated     time.Time `json:"-"`
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
}

//NewStudent Adds a new student to the database
func NewStudent(cName, pinyin, eName, sID, classID, sexID string) (err error) {
	//Create Student
	s := new(Student)
	s.ChineseName = cName
	s.Pinyin = pinyin
	s.EnglishName = eName
	s.StudentID = sID
	s.ClassID, err = strconv.Atoi(classID)
	if err != nil {
		return
	}
	s.SexID, err = strconv.Atoi(sexID)
	if err != nil {
		return
	}

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

	//Add the new student to the database
	conn := db.DB()
	defer conn.Close()
	row := conn.QueryRow(`
		INSERT INTO students 
		(chinese_name, pinyin, english_name,
		 student_id, class_id, sex_id) 
		 values($1, $2, $3, $4, $5, $6)`, s.ChineseName, s.Pinyin, s.EnglishName, s.StudentID, s.ClassID, s.SexID)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(row)
	return
}
