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

//DeleteStudent Sets the delete flag to true
func DeleteStudent(id int) (err error) {
	//Gets a connection to the database
	conn := db.DB()
	defer conn.Close()

	//Determine if student exists
	rows, err := conn.Query(`
	SELECT chinese_name FROM students WHERE id = $1
	`, id)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return
	}

	var name string
	for rows.Next(){
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if strings.EqualFold(name, "") == true {
		err = fmt.Errorf("Student does not exist")
		return
	}

	//Flip the delete flag
	_, err = conn.Query(`UPDATE students SET deleted = true WHERE id = $1`, id)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

//UpdateStudent -- Updates the information for a student
func UpdateStudent(ID, cName, pinyin, eName, sID, classID, sexID string) (err error) {
	//initialize a student
	s := new(Student)
	s.ID, err = strconv.Atoi(ID)
	if err != nil {
		return
	}
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
		//Validation did not pass
		for _, vError := range valid.Errors {
			log.Println(vError.Key, vError.Message)
			err = fmt.Errorf("%s: %s", vError.Key, vError.Message)
			return
		}
	}
	//Update student in the database
	conn := db.DB()
	defer conn.Close()

	_, err = conn.Query(`
	UPDATE students SET
	chinese_name = &2,
	pinyin = &3,
	english_name = &4,
	student_id = &5,
	class_id = &6,
	sex_id = &7
	WHERE id = &1`, s.ChineseName, s.Pinyin, s.EnglishName, s.StudentID, s.ClassID, s.SexID)

	if err != nil {
		log.Println(err)
		return
	}
	return
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
