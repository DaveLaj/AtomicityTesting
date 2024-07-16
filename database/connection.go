package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Test *sql.DB
)

func Connect(creds string, dbname string) error {
	db, err := sql.Open("mysql", creds+"/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	Test = db
	return nil
}
