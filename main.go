package main

import (
	"fmt"
	"log"
	_ "students/routers"
	"students/db"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3" // import your required driver
)

func main() {
	conn := db.DB()

	rows, err := conn.Query("select * from class")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var name string
			err = rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(id, name)
	}
	beego.Run()
}
