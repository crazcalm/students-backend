package models

import (
    "github.com/astaxie/beego/orm"
)

//User -- Learning Model Struct
type User struct {
    ID   int `json:"-"` //"-" omits this field
    Name string `orm:"size(100)", json:"name"`
}

//Student struct to hold student information
type Student struct {
	ID			int		`json:"-"`
	ChineseName string 	`csv:"chinese_name", json:"chinese_name"`
	Pinyin      string 	`csv:"pinyin", json:"pinyin"`
	EnglishName string 	`csv:"english_name", json:english_name`
	StudentID   string 	`csv:"student_id", json:"student_id"`
	Class		string	`json:"class"`
	Sex			string	`json:"-"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Student))
}
