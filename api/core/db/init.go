package db

import (
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const sqlPath = "./sql/transactions.sql"

var (
	globalDB          *sqlx.DB
	globalDBSingleton sync.Once
)

func InitDB(inMemory bool) error {
	var err error

	globalDBSingleton.Do(func() {
		if inMemory {
			globalDB, err = sqlx.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
		} else {
			globalDB, err = sqlx.Open("sqlite3", "./transaction_store.db")
		}

		var sqlBytes []byte
		if sqlBytes, err = os.ReadFile(sqlPath); err != nil {
			return
		}

		if _, err = globalDB.Exec(string(sqlBytes[:])); err != nil {
			return
		}
	})

	return err
}

func GetDB() *sqlx.DB {
	return globalDB
}
