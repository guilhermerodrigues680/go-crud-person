package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type PersonController interface {
	ReadAllHandler(http.ResponseWriter, *http.Request)
	CreateHandler(http.ResponseWriter, *http.Request)
	ReadHandler(http.ResponseWriter, *http.Request)
	UpdateHandler(http.ResponseWriter, *http.Request)
	DeleteHandler(http.ResponseWriter, *http.Request)
}

type PersonRouter struct {
	c PersonController
}

func NewPersonRouter(c PersonController) PersonRouter {
	return PersonRouter{c: c}
}

func (r PersonRouter) Handle(s *mux.Router) {
	s.HandleFunc("", r.c.ReadAllHandler).Methods(http.MethodGet)
	s.HandleFunc("", r.c.CreateHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", r.c.ReadHandler).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", r.c.UpdateHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", r.c.DeleteHandler).Methods(http.MethodDelete)
}
