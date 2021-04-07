package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type RootController interface {
	ApiVersionHandler(http.ResponseWriter, *http.Request)
}

type RootRouter struct {
	c RootController
}

func NewRootRouter(c RootController) RootRouter {
	return RootRouter{c: c}
}

func (r RootRouter) Handle(s *mux.Router) {
	s.HandleFunc("", r.c.ApiVersionHandler).Methods(http.MethodGet)
	s.HandleFunc("/", r.c.ApiVersionHandler).Methods(http.MethodGet)
}
