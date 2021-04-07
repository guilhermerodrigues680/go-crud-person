package main

import (
	"gocrudperson/internal/api"
	"log"
)

const (
	version        = "0.0.1-alpha.0"
	apiAddr string = ":8080"
)

func main() {
	log.Printf("go-crud-person %s (powered by guilhermerodrigues680).", version)
	api.Listen(apiAddr)
}
