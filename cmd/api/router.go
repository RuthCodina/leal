package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	c "github.com/leal/pkg/campaigns/handlers"
	usr "github.com/leal/pkg/user/handlers"
)

type Handlers struct {
	campaign *c.Handler
	user     *usr.Handler
}

func routes(h *Handlers) *chi.Mux {
	log.Println(h)
	r := chi.NewMux()
	r.Get("/", basicHandler)
	r.Post("/", h.campaign.Create)

	r.Get("/get/branch", h.campaign.GetAllByBranch)
	r.Get("/get/commerce", h.campaign.GetAllByCommerce)

	r.Patch("/activate/on", h.campaign.Activate)
	r.Patch("/activate/off", h.campaign.Inactivate)

	r.Patch("/user/accumulate", h.user.AccumulatePoints)

	return r
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("iniciando basic")
	w.Write([]byte("Hello, world!"))
}
