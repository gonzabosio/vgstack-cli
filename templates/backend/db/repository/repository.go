package repository

import "database/sql"

type Service interface {
	LanguageRepository
}

type PostgreService struct {
	DB *sql.DB
}
