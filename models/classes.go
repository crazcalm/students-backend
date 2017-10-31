package models

import (
	_ "github.com/mattn/go-sqlite3" //orm driver
	"log"
	"students/db"
)

//Class -- testing this out
type Class struct {
	ID   int
	Name string
}

//NewClass -- adds a new class to the database
func NewClass(name string) error {
	conn := db.DB()

	stmt, err := conn.Prepare("INSERT INTO class(name) values(?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	if err != nil {
		log.Println(err)
	}
	return err
}
