package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/transaction_test")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	if err := engine.Run(":6900"); err != nil {
		fmt.Println("Error starting server")
	}
}
