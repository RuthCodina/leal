package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/leal/pkg/campaigns/handlers"
)

type Handlers struct {
	campaign handlers.Handler
}

func routes(h *Handlers) *chi.Mux {
	log.Println(h)
	r := chi.NewMux()
	r.Get("/", basicHandler)
	r.Post("/", h.campaign.Create)
	return r
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("iniciando basic")
	w.Write([]byte("Hello, world!"))
}
