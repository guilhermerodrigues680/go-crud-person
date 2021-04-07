package service

import "gocrudperson/internal/model"

// PersonCRUD operações de CRUD para a pessoa
type PersonCRUD interface {
	PersonCreator
	PersonReader
	PersonUpdater
	PersonDeleter
}

// PersonCreator operações de criação de pessoas no sistema
type PersonCreator interface {
	// Cria uma pessoa no sistema
	Create(model.Person) (model.Person, error)
}

// PersonReader operações de leitura/listagem de pessoa no sistema
type PersonReader interface {
	// Read lê uma pessoa no sistema
	Read(int) (model.Person, error)
	// ReadAll lê todas as pessoas no sistema
	ReadAll() ([]model.Person, error)
}

// PersonUpdater operações de edição de pessoa no sistema
type PersonUpdater interface {
	// Update atualiza uma pessoa do sistema
	Update(int, model.Person) (model.Person, error)
}

// PersonDeleter operação de remoção de pessoa do sistema
type PersonDeleter interface {
	// Delete deleta uma pessoa do sistema
	Delete(int) error
}
