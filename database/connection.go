package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DbConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
}

func Connect(config *DbConfig, dbName string) (*sql.DB, error) {
	username := config.DB_USERNAME
	password := config.DB_PASSWORD
	host := config.DB_HOST
	port := config.DB_PORT
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	db, err := sql.Open("mysql", dsn)
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
