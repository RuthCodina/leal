package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

func routes() *chi.Mux {
	r := chi.NewMux()
	r.Get("/", basicHandler)
	return r
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
