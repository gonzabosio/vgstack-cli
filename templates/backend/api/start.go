package api

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func StartBackendServer() error {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	if err := http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), r); err != nil {
		return err
	}
	return nil
}
