package database

import (
	"database/sql"
	"log"
)

var db *sql.DB

func ConnectDB() {
	var err error
	db, err = sql.Open(
		"postgres",
		"user=postgres password=postgres dbname=food_wander sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func DB() *sql.DB {
	return db
}
