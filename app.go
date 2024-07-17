package main

import (
	"SQLTest/config"
	"SQLTest/database"
	"SQLTest/persistence"
	"SQLTest/routes"
	"log"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	db, err := database.Connect("root:@tcp(127.0.0.1:3306)", "transaction_test")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the environment
	env := &config.Env{
		M: persistence.UserModel{DB: db},
	}

	// Initialize the routes
	routes.User(env, engine)

	if err := engine.Run(":6900"); err != nil {
		fmt.Println("Error starting server")
	}
}
