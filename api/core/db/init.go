package db

import (
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	globalDB          *sqlx.DB
	globalDBSingleton sync.Once
)

func InitDB(inMemory bool) error {
	var err error

	globalDBSingleton.Do(func() {
		if inMemory {
			if globalDB, err = sqlx.Open("sqlite3", "file::memory:?mode=memory&cache=shared"); err != nil {
				return
			}
		} else {
			if globalDB, err = sqlx.Open("sqlite3", "./transaction_store.db"); err != nil {
				return
			}
		}

		if _, err = globalDB.Exec(
			`CREATE TABLE IF NOT EXISTS transactions (
				id VARCHAR(32) PRIMARY KEY NOT NULL,
				trans_time TIMESTAMP NOT NULL,
				card_no VARCHAR(16) NOT NULL,
				expiry_month INT NOT NULL,
				expiry_year INT NOT NULL,
				cvv INT NOT NULL,
				currency VARCHAR(3) NOT NULL,
				amount DECIMAL(10, 2) NOT NULL,
				state VARCHAR(32) NOT NULL
			);`,
		); err != nil {
			return
		}
	})

	return err
}

func GetDB() *sqlx.DB {
	return globalDB
}
