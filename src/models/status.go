package models

import (
	"database/sql"
	"log"
	"time"
)

type Status struct {
	Id          int       `json:"id"`
	Status      string    `json:"status"`
	SubLine     string    `json:"subLine"`
	Explanation string    `json:"explanation"`
	Date        time.Time `json:"date"`
}

func setUpStatusesTable(db *sql.DB) {
	statusesQuery := `
		CREATE TABLE IF NOT EXISTS statuses (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
			status TEXT,
			subLine TEXT,
			explanation TEXT,
			date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	_, err := db.Exec(statusesQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *Database) GetMostRecentStatus() Status {
	query := "SELECT * FROM statuses ORDER BY id DESC LIMIT 1"
	row := d.db.QueryRow(query)

	var status Status
	err := row.Scan(&status.Id, &status.Status, &status.SubLine, &status.Explanation, &status.Date)
	if err != nil {
		if err == sql.ErrNoRows {
			return Status{}
		} else {
			log.Fatal(err)
		}
	}
	return status
}

func (d *Database) InsertStatus(status Status) error {
	query := `INSERT INTO statuses (status, subLine, explanation, date) VALUES (?, ?, ?, ?)`
	statement, err := d.db.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(status.Status, status.SubLine, status.Explanation, status.Date)
	if err != nil {
		return err
	}

	return nil
}
