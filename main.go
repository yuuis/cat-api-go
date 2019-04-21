package main

import (
	"github.com/yuuis/cat-api-go/external/api"
	"github.com/yuuis/cat-api-go/inflastructure"
	"github.com/yuuis/cat-api-go/registry"
	"log"
	"os"
)

func main() {
	hash, err := infrastructure.NewHash()
	if err != nil {
		log.Fatal(err)
	}

	db, err := infrastructure.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "cat-api-go", log.Ltime)

	r := registry.NewRegistry(db, hash, logger)
	controller := r.NewController()
	server := infrastructure.NewServer()
	api.NewRouter(server, controller)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
