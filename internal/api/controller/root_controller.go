package controller

import (
	"fmt"
	"gocrudperson/internal/service"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	name := "Joao"

	// person := service.Person{ID: 0, Name: name}

	var service service.PersonCRUD = new(service.PersonService)

	person, err := service.Create(name)

	// personService := new(service.PersonCRUD)

	// personService.

	//person := g.Create(name)

	//person := service.Person.Create(name)

	// person, err := service.Person.Create("ola")

	fmt.Println(person, err)

	w.Write([]byte("API v1 ok!"))
}
