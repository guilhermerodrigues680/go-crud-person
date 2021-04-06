package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func readAllHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("readAllHandler ok!"))
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "readHandler error! Erro: %s\n", err)
		return
	}

	fmt.Fprintf(w, "readHandler ok! ID: %d\n", id)
}

func PersonRouter(s *mux.Router) {
	s.HandleFunc("", readAllHandler)
	s.HandleFunc("/", readAllHandler)
	s.HandleFunc("/{id:[0-9]+}", readHandler)
}
