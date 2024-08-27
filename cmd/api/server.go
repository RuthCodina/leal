package api

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type server struct {
	*http.Server
}

func newServer(listening string, mux *chi.Mux) *server {
	s := &http.Server{
		Addr:         "localhost:" + listening,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &server{s}
}

// To start runs ListendAndServe on the http.server
func (srv *server) Start() {
	log.Println("starting API cmd")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s due to %s", srv.Addr, err.Error())
	}

	//	log.Printf("cmd is ready to handle requests %s", srv.Addr)
}
