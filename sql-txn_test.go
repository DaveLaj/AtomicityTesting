package main_test

import (
	"SQLTest/config"
	"SQLTest/database"
	"SQLTest/persistence"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkTxn(b *testing.B) {

	if err := godotenv.Load(".env"); err != nil {
		b.Errorf("Error loading .env file")
		return
	}

	dbconf := &database.DbConfig{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
	}

	db, err := database.Connect(dbconf, "gk_miniapps")
	if err != nil {
		b.Errorf("Error connecting to database: %v", err)
		return
	}

	b.Run("TestTxn", func(t *testing.B) {

		env := &config.Env{
			Model: persistence.UserModel{DB: db},
		}
		id := 2

		UserModel := env.Model

		txnDone := make(chan bool)
		go func() {
			err = UserModel.UpdateAmountByIDTxn(id, 69)
			if err != nil {
				t.Errorf("Error while updating row: %v", err)
				return
			}
			txnDone <- true
		}()

		err = UserModel.UpdateAmountByIDTxn(id+1, 69)
		if err != nil {
			t.Errorf("Error while updating row: %v", err)
			return
		}

		<-txnDone
	})
}
