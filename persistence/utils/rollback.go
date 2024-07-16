package utils

import (
	"database/sql"
)

func Rollback(tx *sql.Tx, HasCommit bool) error {
	if !HasCommit {
		err := tx.Rollback()
		if err != nil {
			return err
		}
	}
	return nil
}
