package store

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const sqlRelativePath = "./sql/transactions"

var (
	globalDB          *sqlx.DB
	globalDBSingleton sync.Once
)

func InitDB() error {
	var err error

	globalDBSingleton.Do(func() {
		globalDB, err = sqlx.Open("sqlite3", "./transaction_store.db")

		var sqlBytes []byte
		sqlItems, _ := os.ReadDir(sqlRelativePath)
		for i := 0; i < len(sqlItems); i++ {
			if !sqlItems[i].IsDir() {
				sqlBytes, err = os.ReadFile(path.Join(sqlRelativePath, sqlItems[i].Name()))
				if err != nil {
					return
				}

				if _, err = globalDB.Exec(string(sqlBytes[:])); err != nil {
					return
				}
				log.Printf("INFO: Executed %s", sqlItems[i].Name())
			} else {
				log.Printf("WARN: Skipping directory %s", sqlItems[i].Name())
			}
		}
	})

	return err
}

func GetDB() *sqlx.DB {
	return globalDB
}
