package repository

import (
	"fmt"
	"gocrudperson/internal/model"
)

type PersonRepositoryLocal struct{}

var persons []model.Person
var pk int64 = 0

func (PersonRepositoryLocal) Insert(person model.Person) (model.Person, error) {
	pk++
	person.ID = pk
	persons = append(persons, person)
	fmt.Println(persons)
	return person, nil
}
