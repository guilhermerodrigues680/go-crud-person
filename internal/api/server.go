package api

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"gocrudperson/internal/api/controller"
	"gocrudperson/internal/api/middleware"
	"gocrudperson/internal/api/route"
	"gocrudperson/internal/repository"
	"gocrudperson/internal/service"

	"github.com/gorilla/mux"
)

type Server struct{}

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

func getPersonRouter() route.PersonRouter {
	r := repository.NewPersonRepositoryLocal()
	s := service.NewPersonService(r)
	c := controller.NewPersonController(s)
	return route.NewPersonRouter(c)
}

func getRootRouter() route.RootRouter {
	c := controller.NewRootController()
	return route.NewRootRouter(c)
}

func NewServer() Server {
	return Server{}
}

func (Server) Listen(addr string) {
	r := mux.NewRouter()

	r.Use(middleware.NewLoggingMiddleware().MiddlewareFunc)
	getRootRouter().Handle(r.PathPrefix("/api/v1").Subrouter())
	getPersonRouter().Handle(r.PathPrefix("/api/v1/person").Subrouter())

	dirWeb, dirOpenApi, err := getWebAndOpenApiDir()
	if err != nil {
		panic(err)
	}

	// Static Files
	r.PathPrefix("/api/openapi-spec/").Handler(http.StripPrefix("/api/openapi-spec/", http.FileServer(http.Dir(dirOpenApi))))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(dirWeb)))

	log.Printf("Web static dir: %s\n", dirWeb)
	log.Printf("OpenAPI static dir: %s\n", dirOpenApi)
	log.Printf("Listening on: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
