package models

import (
    _ "github.com/mattn/go-sqlite3" //orm driver
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
func NewUser(name string) {
	user := new(User)
	user.Name = name
}
