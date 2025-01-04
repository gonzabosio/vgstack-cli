package api

import (
	"net/http"
	"vgstack-cli/templates/backend/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func newRouter() (*chi.Mux, error) {
	h, err := handlers.NewHandler()
	if err != nil {
		return nil, err
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a programming languages API"))
	})
	r.Route("/language", func(r chi.Router) {
		r.Post("/", h.AddLanguage)
		r.Delete("/{lang_id}", h.DeleteLanguage)
	})
	return r, nil
}
