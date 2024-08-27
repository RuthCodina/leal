package main

import (
	"os"

	"github.com/leal/cmd/api"
)

const defaultPort = "8000"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	api.Start(port)
}
