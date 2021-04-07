package controller

import (
	"encoding/json"
	"fmt"
	"gocrudperson/internal/model"
	"gocrudperson/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var personService service.PersonCRUD = new(service.PersonService)

func genericErrorHandler(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Erro: %s", err)
}

func readAllHandler(w http.ResponseWriter, r *http.Request) {
	persons, err := personService.ReadAll()
	if err != nil {
		genericErrorHandler(w, err)
		return
	}

	mapObj := make(map[string][]model.Person)
	mapObj["persons"] = persons
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mapObj)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		genericErrorHandler(w, err)
		return
	}

	person, err := personService.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Erro: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		genericErrorHandler(w, err)
		return
	}

	err = personService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Erro: %s", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	personBody := new(model.Person)
	err := json.NewDecoder(r.Body).Decode(personBody)
	if err != nil {
		genericErrorHandler(w, err)
		return
	}

	person, err := personService.Create(*personBody)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	personBody := new(model.Person)
	err := json.NewDecoder(r.Body).Decode(personBody)
	if err != nil {
		genericErrorHandler(w, err)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		genericErrorHandler(w, err)
		return
	}

	person, err := personService.Update(id, *personBody)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Erro: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func PersonRouter(s *mux.Router) {
	s.HandleFunc("", readAllHandler).Methods(http.MethodGet)
	s.HandleFunc("", createHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id:[0-9]+}", readHandler).Methods(http.MethodGet)
	s.HandleFunc("/{id:[0-9]+}", updateHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id:[0-9]+}", deleteHandler).Methods(http.MethodDelete)
}
