package persistence

import (
	"database/sql"
	"fmt"

	"SQLTest/models"
)

// SelectRow is a function to select a row from a table
func (m *UserModel) SelectRowByIDTxn(id int) (user *models.User, err error) {
	stmt := `SELECT id, name, age FROM test WHERE id = ?`
	test := m.DB
	tx, err := test.Begin()
	if err != nil {
		return nil, fmt.Errorf("Error while starting transaction: %w", err)
	}
	user = &models.User{}
	row := tx.QueryRow(stmt, id)

	err = row.Scan(&user.ID, &user.Name, &user.Age)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("No rows found")
	}
	if err != nil {
		return nil, fmt.Errorf("Error while scanning row: %w", err)
	}

	return user, nil
}
