package utils

import "database/sql"

// begin is a function to begin a transaction
func begin(db *sql.DB) (*sql.Tx, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}
