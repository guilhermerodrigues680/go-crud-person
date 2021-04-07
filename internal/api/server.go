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

func getWebAndOpenApiDir() (string, string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", "", err
	}

	exPath := filepath.Dir(ex)
	dirWeb, err := filepath.Abs(exPath + "/../web")
	if err != nil {
		return "", "", err
	}

	dirOpenApi, err := filepath.Abs(exPath + "/../api/openapi-spec")
	if err != nil {
		return "", "", err
	}

	return dirWeb, dirOpenApi, nil
}

func Listen(addr string) {
	r := mux.NewRouter()

	dirWeb, dirOpenApi, err := getWebAndOpenApiDir()
	if err != nil {
		panic(err)
	}

	r.Use(middleware.LoggingMiddleware)
	r.HandleFunc("/api/v1", controller.RootHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/", controller.RootHandler).Methods(http.MethodGet)
	controller.PersonRouter(r.PathPrefix("/api/v1/person").Subrouter())

	// Static Files
	r.PathPrefix("/api/openapi-spec/").Handler(http.StripPrefix("/api/openapi-spec/", http.FileServer(http.Dir(dirOpenApi))))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(dirWeb)))

	log.Printf("Web static dir: %s\n", dirWeb)
	log.Printf("OpenAPI static dir: %s\n", dirOpenApi)
	log.Printf("Listening on: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
