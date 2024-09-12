package db

import (
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// Using a global singleton as assuming a single database and as sqlx.DB already providers connection pooling.
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

		if err = initTables(globalDB); err != nil {
			return
		}
	})

	return err
}

func initTables(aDB *sqlx.DB) error {
	return createTransactionsTable(aDB)
}

func GetDB() *sqlx.DB {
	return globalDB
}
