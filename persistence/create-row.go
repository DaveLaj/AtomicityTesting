package persistence

import (
	"SQLTest/models"
	utils "SQLTest/persistence/utils"
	"database/sql"
	"fmt"
)

type UserModel struct {
	DB *sql.DB
}

// CreateRow is a function to create a row in a table
func (m *UserModel) CreateRow(user models.CreateUser) error {
	test := m.DB
	tx, err := test.Begin()
	if err != nil {
		return err
	}
	var HasCommit bool = false
	defer func() {
		err = utils.Rollback(tx, HasCommit)
		if err != nil {
			// better to use panic here to recover from the error
			// using a goroutine that recovers when database conn is working
			panic(fmt.Errorf("Rollback Error!: %w", err))
		}
	}()

	_, err = tx.Exec("INSERT INTO test (name, age) VALUES (?, ?)", user.Name, user.Age)
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE test SET age = age + 100 WHERE name = ?", user.Name)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	HasCommit = true
	return nil
}
