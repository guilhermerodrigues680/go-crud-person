package service

import (
	"gocrudperson/internal/model"
	"gocrudperson/internal/repository"
)

type PersonService struct{}

var personRepository repository.PersonRepository = new(repository.PersonRepositoryLocal)

func (s *PersonService) Create(name string) (model.Person, error) {
	p := model.Person{ID: -1, Name: name}
	p, _ = personRepository.Insert(p)
	return p, nil
}
