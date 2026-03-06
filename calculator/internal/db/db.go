package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func Connect() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to DB succesfully!")

	err = initTable(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func initTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS calc_history (
		id BIGSERIAL PRIMARY KEY,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		mode TEXT NOT NULL,
		input JSONB NOT NULL,
		success BOOLEAN NOT NULL,
		output JSONB,
		error TEXT,
		duration_ms BIGINT,
		note TEXT
	);
	`)
	return err
}
