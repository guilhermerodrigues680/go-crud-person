package repository

import "gocrudperson/internal/model"

// PersonCRUD operações de CRUD para a pessoa
type PersonRepository interface {
	Insert(model.Person) (model.Person, error)
}
