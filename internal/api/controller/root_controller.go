package controller

import (
	"fmt"
	"gocrudperson/internal/config"
	"net/http"
)

type RootController struct{}

func NewRootController() RootController {
	return RootController{}
}

func (RootController) ApiVersionHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("API v1 ok!"))
	conf := config.GetConfig()
	fmt.Fprintf(w, "%s %s (powered by guilhermerodrigues680)", conf.AppName, conf.Version)
}
