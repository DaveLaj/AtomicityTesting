package persistence

import (
	"fmt"
)

// SelectRow is a function to select a row from a table
func (m *UserModel) UpdateNameByIDTxn(id int, name string) error {
	stmt := `UPDATE test SET name = ? WHERE id = ?`
	test := m.DB

	tx, err := test.Begin()
	if err != nil {
		return fmt.Errorf("Error while starting transaction: %w", err)
	}

	_, err = tx.Exec(stmt, name, id)
	if err != nil {
		return fmt.Errorf("Error while updating row: %w", err)
	}

	// time.Sleep(5 * time.Second)

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Error while committing transaction: %w", err)
	}

	return nil
}
