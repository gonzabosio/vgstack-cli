package model

type Language struct {
	Id          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	ReleaseYear int    `json:"release_year" validate:"required"`
}
