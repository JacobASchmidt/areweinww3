package models

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

var db Database
var once sync.Once

func GetInstance() Database {
	once.Do(func() {
		database, err := sql.Open("sqlite3", "./test.db")
		if err != nil {
			log.Fatal(err)
		}

		setUpArticlesTable(database)
		setUpStatusesTable(database)
		db = Database{db: database}
	})

	return db
}

func (db *Database) Close() {
	db.db.Close()
}
