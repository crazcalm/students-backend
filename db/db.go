package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq" //database driver
	"io/ioutil"
	"log"
)

func getConfig() map[string]string {
	var config map[string]string

	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

//DB gives out connections to the database
func DB() *sql.DB {
	//Get secret values from config
	config := getConfig()

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s", config["db_user"], config["db_name"], config["db_user_password"])
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
