package handlers

import (
	"github.com/gonzabosio/vgstack-cli/templates/backend/db"
	"github.com/gonzabosio/vgstack-cli/templates/backend/db/repository"
)

type Handler struct {
	rp repository.Service
}

func NewHandler() (*Handler, error) {
	db, err := db.NewDB()
	if err != nil {
		return nil, err
	}
	return &Handler{rp: &repository.PostgreService{DB: db}}, nil
}
