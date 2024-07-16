package main

import (
	"SQLTest/database"
	"SQLTest/routes"
	"database/sql"

	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Rollback(tx *sql.Tx) {
	if tx == nil {
		return
	}
	if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
		log.Fatal(err)
	}
}

func AddEntry(db *sql.DB, name string, age int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer Rollback(tx)

	_, err = tx.Exec("INSERT INTO test (name, age) VALUES (?, ?)", name, age)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE test SET age = age + 100 WHERE name = ?", name)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func main() {

	err := database.Connect("root:@tcp(127.0.0.1:3306)", "transaction_test")
	if err != nil {
		log.Fatal(err)
	}
	engine := gin.Default()

	routes.User(engine)
	// engine.GET("/", func(c *gin.Context) {
	// 	err := AddEntry(db, "Johnny Joestar", 20)
	// 	if err != nil {
	// 		c.JSON(500, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(200, gin.H{
	// 		"message": "Transaction is Successful!",
	// 	})
	// })

	if err := engine.Run(":6900"); err != nil {
		fmt.Println("Error starting server")
	}
}
