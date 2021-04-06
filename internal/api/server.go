package api

import (
	"log"
	"net/http"

	"gocrudperson/internal/api/controller"
	"gocrudperson/internal/api/middleware"

	"github.com/gorilla/mux"
)

func Listen(addr string) {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/api/v1", controller.RootHandler)
	controller.PersonRouter(r.PathPrefix("/api/v1/person").Subrouter())

	log.Println("go-crud-person. Listening on: " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
