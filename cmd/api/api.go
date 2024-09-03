package api

import (
	"github.com/leal/db"
	ch "github.com/leal/pkg/campaigns/handlers"
	cr "github.com/leal/pkg/campaigns/repository"
	cs "github.com/leal/pkg/campaigns/services"
	uh "github.com/leal/pkg/users/handlers"
	ur "github.com/leal/pkg/users/repository"
	us "github.com/leal/pkg/users/services"
)

func Start(port string) {
	db := db.ConnectToDB()
	campaignRepo := cr.NewCampaingRepository(db)
	campaignsrv := cs.NewCampaingService(campaignRepo)
	campaignHandler := ch.NewCampaignHandler(campaignsrv)

	userRepo := ur.NewUserRepository(db)
	usersrv := us.NewUserService(userRepo, campaignsrv)
	userHandler := uh.NewUserHandler(usersrv)

	h := Handlers{
		campaign: campaignHandler,
		user:     userHandler,
	}
	r := routes(&h)
	server := newServer(port, r)
	server.Start()
}
