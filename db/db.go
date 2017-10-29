package db

import(
	"os"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3" //database driver
)

//DB gives out connections to the database
func DB() *sql.Tx {
	os.Remove("./testing.db")
	
	db, err := sql.Open("sqlite3", "./db/testing.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
		PRAGMA foreign_keys = ON;
		`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	
	return tx
}

