package api

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gocrudperson/internal/api/controller"
	"gocrudperson/internal/api/middleware"

	"github.com/gorilla/mux"
)

func getWebDir() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	exPath := filepath.Dir(ex)
	dir, err := filepath.Abs(exPath + "/../web")
	if err != nil {
		return "", err
	}

	return dir, nil
}

func Listen(addr string) {
	r := mux.NewRouter()

	dir, err := getWebDir()
	if err != nil {
		panic(err)
	}

	r.Use(middleware.LoggingMiddleware)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))
	r.HandleFunc("/api/v1", controller.RootHandler).Methods(http.MethodGet)
	controller.PersonRouter(r.PathPrefix("/api/v1/person").Subrouter())

	log.Printf("Web static dir: %s\n", dir)
	log.Printf("Listening on: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
