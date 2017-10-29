package models

import (
	_ "github.com/mattn/go-sqlite3" //orm driver
)

//Class -- testing this out
type Class struct {
	ID		int
	Name	string
}
