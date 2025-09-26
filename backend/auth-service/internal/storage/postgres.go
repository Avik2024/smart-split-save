package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// NewPostgres connects to PostgreSQL and returns the DB instance
func NewPostgres(dbURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Migrate creates required tables
func Migrate(db *sqlx.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL,
		name TEXT,
		created_at TIMESTAMP DEFAULT NOW()
	);`
	_, err := db.Exec(schema)
	if err != nil {
		log.Println("Error running migrations:", err)
		return err
	}
	return nil
}
