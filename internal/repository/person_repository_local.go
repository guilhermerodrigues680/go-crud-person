package repository

import (
	"errors"
	"fmt"
	"gocrudperson/internal/model"
)

type PersonRepositoryLocal struct {
	persons []model.Person
	pk      int
}

func NewPersonRepositoryLocal() PersonRepositoryLocal {
	return PersonRepositoryLocal{
		persons: make([]model.Person, 0),
		pk:      0,
	}
}

func (r PersonRepositoryLocal) Insert(newPerson model.Person) (model.Person, error) {
	r.pk++
	newPerson.ID = r.pk
	r.persons = append(r.persons, newPerson)
	fmt.Println(r.persons)
	return newPerson, nil
}

func (r PersonRepositoryLocal) SelectAll() ([]model.Person, error) {
	return r.persons, nil
}

func (r PersonRepositoryLocal) Select(id int) (model.Person, error) {
	for _, person := range r.persons {
		if person.ID == id {
			return person, nil
		}
	}
	return model.Person{}, errors.New("Não há uma pessoa cadastrada com o id")
}

func (r PersonRepositoryLocal) Update(id int, newPerson model.Person) (model.Person, error) {
	for index, person := range r.persons {
		if person.ID == id {
			r.persons = append(r.persons[:index], r.persons[index+1:]...)
			newPerson.ID = id
			r.persons = append(r.persons, newPerson)
			return newPerson, nil
		}
	}
	return model.Person{}, errors.New("Não há uma pessoa cadastrada com o id")
}

func (r PersonRepositoryLocal) Delete(id int) error {
	for index, person := range r.persons {
		if person.ID == id {
			r.persons = append(r.persons[:index], r.persons[index+1:]...)
			return nil
		}
	}
	return errors.New("Não há uma pessoa cadastrada com o id")
}
