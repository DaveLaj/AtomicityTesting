package persistence

import (
	"fmt"
	"time"
)

// SelectRow is a function to select a row from a table
func (m *UserModel) UpdateAmountByIDTxn(id int, amount int) error {
	stmt := `UPDATE wallet_logs SET Amount = ? WHERE id = ?`
	test := m.DB

	tx, err := test.Begin()
	if err != nil {
		return fmt.Errorf("Error while starting transaction: %w", err)
	}

	_, err = tx.Exec(stmt, amount, id)
	if err != nil {
		return fmt.Errorf("Error while updating row: %w", err)
	}

	time.Sleep(5 * time.Second)

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("Error while committing transaction: %w", err)
	}

	return nil
}
