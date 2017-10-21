package models

import (
    "github.com/astaxie/beego/orm"
)

//User -- Learning Model Struct
type User struct {
    ID   int `json:"-"` //"-" omits this field
    Name string `orm:"size(100)", json:"name"`
}

func init() {
	// register model
	orm.RegisterModel(new(User))
}
