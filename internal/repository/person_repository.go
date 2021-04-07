package repository

import "gocrudperson/internal/model"

// PersonRepository operações de CRUD no banco de dados para a pessoa
type PersonRepository interface {
	// Insert insere uma pessoa no banco de dados
	Insert(model.Person) (model.Person, error)
	// SelectAll seleciona todas as pessoas do banco de dados
	SelectAll() ([]model.Person, error)
	// Select seleciona uma pessoa do banco de dados
	Select(int) (model.Person, error)
	// Update atualiza uma pessoa no banco de dados
	Update(int, model.Person) (model.Person, error)
	// Delete apaga uma pessoa do banco banco de dados
	Delete(int) error
}
