package main

import (
	"gocrudperson/internal/api"
	"gocrudperson/internal/config"
	"log"
)

func main() {
	conf := config.GetConfig()
	log.Printf("%s %s (powered by guilhermerodrigues680)", conf.AppName, conf.Version)
	api.NewServer().Listen(conf.ServerAddr)
}
