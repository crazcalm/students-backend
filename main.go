package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq" // import your required driver
	"log"
	"students/db"
	_ "students/routers"
)

func main() {
	conn := db.DB()

	rows, err := conn.Query("select id, name from class")
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
