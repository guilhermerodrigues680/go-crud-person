package repository

import (
	"errors"
	"fmt"
	"gocrudperson/internal/model"
)

type PersonRepositoryLocal struct{}

var persons []model.Person = make([]model.Person, 0)
var pk int = 0

func (PersonRepositoryLocal) Insert(newPerson model.Person) (model.Person, error) {
	pk++
	newPerson.ID = pk
	persons = append(persons, newPerson)
	fmt.Println(persons)
	return newPerson, nil
}

func (PersonRepositoryLocal) SelectAll() ([]model.Person, error) {
	return persons, nil
}

func (PersonRepositoryLocal) Select(id int) (model.Person, error) {
	for _, person := range persons {
		if person.ID == id {
			return person, nil
		}
	}
	return model.Person{}, errors.New("Não há uma pessoa cadastrada com o id")
}

func (PersonRepositoryLocal) Update(id int, newPerson model.Person) (model.Person, error) {
	for index, person := range persons {
		if person.ID == id {
			persons = append(persons[:index], persons[index+1:]...)
			newPerson.ID = id
			persons = append(persons, newPerson)
			return newPerson, nil
		}
	}
	return model.Person{}, errors.New("Não há uma pessoa cadastrada com o id")
}

func (PersonRepositoryLocal) Delete(id int) error {
	for index, person := range persons {
		if person.ID == id {
			persons = append(persons[:index], persons[index+1:]...)
			return nil
		}
	}
	return errors.New("Não há uma pessoa cadastrada com o id")
}
