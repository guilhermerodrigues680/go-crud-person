package controller

import (
	"encoding/json"
	"fmt"
	"gocrudperson/internal/model"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// PersonCRUD operações de CRUD para a pessoa
type PersonCRUD interface {
	Create(model.Person) (model.Person, error)
	Read(int) (model.Person, error)
	ReadAll() ([]model.Person, error)
	Update(int, model.Person) (model.Person, error)
	Delete(int) error
}

type Controller struct {
	s PersonCRUD
}

func NewPersonController(s PersonCRUD) Controller {
	return Controller{s: s}
}

func genericErrorHandler(w *http.ResponseWriter, r *http.Request, err error) {
	type GenericError struct {
		Timestamp time.Time `json:"timestamp"`
		Status    int       `json:"status"`
		Error     string    `json:"error"`
		Message   string    `json:"message"`
		Path      string    `json:"path"`
	}

	log.Printf("genericErrorHandler - Erro: %s", err)

	genericError := GenericError{
		Timestamp: time.Now().UTC(),
		Status:    http.StatusInternalServerError,
		Error:     http.StatusText(http.StatusInternalServerError),
		Message:   err.Error(),
		Path:      r.RequestURI,
	}

	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(*w).Encode(genericError)
}

func (c Controller) ReadAllHandler(w http.ResponseWriter, r *http.Request) {
	type PersonResponse struct {
		Persons []model.Person `json:"persons"`
	}

	persons, err := c.s.ReadAll()
	if err != nil {
		genericErrorHandler(&w, r, err)
		return
	}

	personResponse := PersonResponse{Persons: persons}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(personResponse)
}

func (c Controller) ReadHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		genericErrorHandler(&w, r, err)
		return
	}

	person, err := c.s.Read(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Erro: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (c Controller) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		genericErrorHandler(&w, r, err)
		return
	}

	err = c.s.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Erro: %s", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c Controller) CreateHandler(w http.ResponseWriter, r *http.Request) {
	personBody := new(model.Person)
	err := json.NewDecoder(r.Body).Decode(personBody)
	if err != nil {
		genericErrorHandler(&w, r, err)
		return
	}

	person, err := c.s.Create(*personBody)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (c Controller) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	personBody := new(model.Person)
	err := json.NewDecoder(r.Body).Decode(personBody)
	if err != nil {
		genericErrorHandler(&w, r, err)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		genericErrorHandler(&w, r, err)
		return
	}

	person, err := c.s.Update(id, *personBody)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Erro: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}
