package main

import (
	"log"
	"net/http"

	"gocrudperson/internal/api/v1/controller"

	"github.com/gorilla/mux"
)

const host string = ":8080"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1", controller.RootHandler)

	log.Println("go-crud-person. Listening on: " + host)
	http.ListenAndServe(host, r)
}
