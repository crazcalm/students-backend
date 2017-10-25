package models

import (
    "github.com/astaxie/beego/orm"
)

//User -- Learning Model Struct
type User struct {
    ID   int `json:"-", orm:"auto"`
    Name string `orm:"size(100)", json:"name"`
}

//TableName -- defining the table name
func (u *User) TableName() string {
    // db table name
    return "user"
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
	orm.RegisterModel(new(User))
}
