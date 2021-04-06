package service

import "gocrudperson/internal/model"

// // Person representa uma pessoa no sistema
// type Person struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name"`
// }

// PersonCRUD operações de CRUD para a pessoa
// type PersonCRUD interface {
// 	// Create(string) (Person, error)
// 	// Create(string) Person
// }

// PersonCRUD operações de CRUD para a pessoa
type PersonCRUD interface {
	PersonCreator
	// PersonReader
	// PersonUpdater
	// PersonDeleter
}

// PersonCreator operações de criação de pessoas no sistema
type PersonCreator interface {
	Create(string) (model.Person, error)
}

// // PersonReader operações de leitura/listagem de pessoa no sistema
// type PersonReader interface {
// 	// Read lê uma pessoa no sistema
// 	Read(int64) (Person, error)
// 	// ReadAll lê todas as pessoas no sistema
// 	ReadAll() ([]Person, error)
// }

// // PersonUpdater operações de edição de pessoa no sistema
// type PersonUpdater interface {
// 	// Update atualiza uma pessoa do sistema
// 	Update(int64, Person) (Person, error)
// }

// // PersonDeleter operação de remoção de pessoa do sistema
// type PersonDeleter interface {
// 	// Delete deleta uma pessoa do sistema
// 	Delete(int64) error
// }
