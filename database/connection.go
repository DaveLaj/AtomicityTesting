package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(creds string, dbname string) (*sql.DB, error) {
	db, err := sql.Open("mysql", creds+"/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

// This is for when you have multiple databases to connect to
// type DatabasePools struct {
// 	db [1]*sql.DB
// }

// func InitializeDatabasePools() *DatabasePools {
// 	var dbPools [1]*sql.DB
// 	db, err := Connect("root:@tcp(127.0.0.1:3306)", "transaction_test")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	dbPools[0] = db
// 	return &DatabasePools{db: dbPools}
// }
