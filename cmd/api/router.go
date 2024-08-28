package api

import (
	"net/http"

	"github.com/go-chi/chi"
	campaigns "github.com/leal/pkg/campaigns/service"
)

func routes() *chi.Mux {
	compaignsrv := campaigns.Service{}

	r := chi.NewMux()
	r.Get("/", basicHandler)
	//	r.Post("/campaign",  )
	return r
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
