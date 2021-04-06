package main

import (
	"log"
	"net/http"

	"gocrudperson/internal/api/v1/controller"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Executando app")
	r := mux.NewRouter()

	r.HandleFunc("/api/v1", controller.RootHandler)

	http.ListenAndServe(":8080", r)
}
