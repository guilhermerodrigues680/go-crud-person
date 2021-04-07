package service

import (
	"gocrudperson/internal/model"
	"gocrudperson/internal/repository"
)

type PersonService struct{}

var personRepository repository.PersonRepository = new(repository.PersonRepositoryLocal)

func (PersonService) Create(newPerson model.Person) (model.Person, error) {
	return personRepository.Insert(newPerson)
}

func (PersonService) Read(id int) (model.Person, error) {
	return personRepository.Select(id)
}

func (PersonService) ReadAll() ([]model.Person, error) {
	return personRepository.SelectAll()
}

func (PersonService) Update(id int, newPerson model.Person) (model.Person, error) {
	return personRepository.Update(id, newPerson)
}

func (PersonService) Delete(id int) error {
	return personRepository.Delete(id)
}
