package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	HOST = "database"
	PORT = 5432
)

// ErrNoMatch is returned when we request a row that doesn't exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")

	migrationQuery, err1 := os.ReadFile("./db/migrations/000001_create_items_table.up.sql")
	if err1 != nil {
		log.Println(err1)
		return db, err
	}
	log.Println("Running migration...")
	log.Println(string(migrationQuery))
	db.Conn.QueryRow(string(migrationQuery))
	log.Println("Migrations done")
	return db, nil
}
