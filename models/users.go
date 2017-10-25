package models

import (
    "github.com/astaxie/beego/orm"
)

//User -- Learning Model Struct
type User struct {
    ID   int `json:"-"`
    Name string `orm:"size(100)", json:"name"`
}

//TableName -- defining the table name
func (u *User) TableName() string {
    // db table name
    return "user"
}

//Student struct to hold student information
type Student struct {
	ID			int		`json:"-"`
	ChineseName string 	`csv:"chinese_name", json:"chinese_name"`
	Pinyin      string 	`csv:"pinyin", json:"pinyin"`
	EnglishName string 	`csv:"english_name", json:english_name`
	StudentID   string 	`csv:"student_id", json:"student_id", orm:"student_id"`
	Class		string	`json:"class"`
	Sex			string	`json:"-"`
}

//NewUser -- testing this out
func NewUser(name string) error {
	o := orm.NewOrm()
	user := new(User)
	user.Name = name
	_,err := o.Insert(user)
	return err
}

func init() {
	// register model
	orm.RegisterModel(new(User), new(Student))
}
