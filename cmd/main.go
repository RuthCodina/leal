package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leal/cmd/api"
	"github.com/leal/db"
)

const defaultPort = "8000"

func main() {
	// connect to DB
	db.ConnectToDB()

	// start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	api.Start(port)

}
