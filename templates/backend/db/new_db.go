package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_CONN_STR"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
