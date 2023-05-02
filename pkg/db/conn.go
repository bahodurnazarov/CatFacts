package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // add this
)

func ConnDB() *sql.DB {
	connStr := "postgresql://postgres:postgres@localhost/catFacts?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
