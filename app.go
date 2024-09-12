package main

import (
	"SQLTest/config"
	"SQLTest/database"
	"SQLTest/persistence"
	"SQLTest/routes"
	"log"
	"os"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	engine := gin.Default()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConfig := &database.DbConfig{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
	}

	db, err := database.Connect(dbConfig, "gk_miniapps")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the environment
	env := &config.Env{
		Model: persistence.UserModel{DB: db},
	}

	// Initialize the routes
	routes.User(env, engine)

	if err := engine.Run(":6900"); err != nil {
		fmt.Println("Error starting server")
	}
}
