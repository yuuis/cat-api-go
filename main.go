package main

import (
	"github.com/yuuis/cat-api-go/external/api"
	"github.com/yuuis/cat-api-go/infrastructure"
	"github.com/yuuis/cat-api-go/registry"
	"log"
	"os"
)

func main() {
	db, err := infrastructure.OpenDB()
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(os.Stdout, "cat-api-go", log.Ltime)

	r := registry.NewRegistry(db, logger)
	c := r.NewController()
	s := infrastructure.NewServer()
	api.NewRouter(s, c)

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
