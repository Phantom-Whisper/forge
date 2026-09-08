package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/database/sqlite3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration command (up, down, or reset)")
	}

	command := os.Args[1]

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	instance, err := sqlite3.WithInstance(db, &sqlite3.Config{})
}