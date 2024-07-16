package persistence

import (
	"database/sql"
	"fmt"

	"SQLTest/database"
	"SQLTest/models"
)

// SelectRow is a function to select a row from a table
func SelectRow(id int) (user *models.User, err error) {
	stmt := `SELECT id, name, age FROM test WHERE id = ?`
	test := database.Test
	user = &models.User{}
	row := test.QueryRow(stmt, id)

	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("No rows found")
	}
	if err != nil {
		return nil, fmt.Errorf("Error while scanning row: %w", err)
	}

	return user, nil
}
