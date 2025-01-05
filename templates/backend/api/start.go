package api

import "github.com/go-chi/chi/v5"

func SetupBackendServer() (*chi.Mux, error) {
	r, err := newRouter()
	if err != nil {
		return nil, err
	}

	return r, nil
}
