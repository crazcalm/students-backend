package models

import (
    "github.com/astaxie/beego/orm"
)

//Student struct to hold student information
type Student struct {
	ID			int		`json:"-", orm:"auto"`
	ChineseName string 	`csv:"chinese_name", json:"chinese_name"`
	Pinyin      string 	`csv:"pinyin", json:"pinyin"`
	EnglishName string 	`csv:"english_name", json:english_name`
	StudentID   string 	`csv:"student_id", json:"student_id", orm:"student_id"`
	Class		string	`json:"class"`
	Sex			string	`json:"-"`
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

	//Insert into the database
	_, err := o.Insert(s)
	return err
}

func init() {
	// register model
	orm.RegisterModel(new(Student))
}
