package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:./data/sqlite3.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
	}

	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Fatalf("Could not create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./sql/migrations",
		"sqlite3", driver)
	if err != nil {
		log.Fatalf("Could not create migrate instance: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration command: up, down, force")
	}

	switch os.Args[1] {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Could not apply migrations: %v", err)
		}
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Could not rollback migrations: %v", err)
		}
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a version to force")
		}
		version := os.Args[2]
		v, err := strconv.Atoi(version)
		if err != nil {
			log.Fatalf("Invalid version: %v", err)
		}
		if err := m.Force(v); err != nil {
			log.Fatalf("Could not force version: %v", err)
		}
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}

	log.Println("Migration complete")
}
