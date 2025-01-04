package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func execQuery(db *sql.DB, queryPath string) error {
	q, err := os.ReadFile(queryPath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(q))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := godotenv.Load("backend/.env"); err != nil {
		log.Fatalf("Failed to load enviroment variables: %v", err)
	}
	db, err := sql.Open("postgres", os.Getenv("DB_CONN_STR"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	basePath := "backend/db/migration/query"

	if err := execQuery(db, basePath+"/create_table_lang.sql"); err != nil {
		log.Fatalf("Unable to create/verify language table: %v", err)
	}
	log.Println("Database migration/update successful")
}
