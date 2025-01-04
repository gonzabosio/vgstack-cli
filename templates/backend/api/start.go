package api

import (
	"net/http"
	"os"
)

func StartBackendServer() error {
	r, err := newRouter()
	if err != nil {
		return err
	}
	if err := http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), r); err != nil {
		return err
	}
	return nil
}
