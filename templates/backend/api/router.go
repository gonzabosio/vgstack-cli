package api

import (
	"net/http"
	"os"
	"vgstack-cli/templates/backend/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func newRouter() (*chi.Mux, error) {
	h, err := handlers.NewHandler()
	if err != nil {
		return nil, err
	}
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{os.Getenv("FRONTEND_URL")},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
	}))
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is a programming languages API"))
	})
	r.Route("/language", func(r chi.Router) {
		r.Post("/", h.AddLanguage)
		r.Get("/", h.GetLanguages)
		r.Delete("/{lang_id}", h.DeleteLanguage)
	})
	return r, nil
}
