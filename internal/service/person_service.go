package service

import "gocrudperson/internal/model"

// PersonRepository representa um repositório de pessoas
type PersonRepository interface {
	Insert(model.Person) (model.Person, error)
	Select(int) (model.Person, error)
	SelectAll() ([]model.Person, error)
	Update(int, model.Person) (model.Person, error)
	Delete(int) error
}

type PersonService struct {
	r PersonRepository
}

func NewPersonService(r PersonRepository) PersonService {
	return PersonService{r: r}
}

func (s PersonService) Create(newPerson model.Person) (model.Person, error) {
	return s.r.Insert(newPerson)
}

func (s PersonService) Read(id int) (model.Person, error) {
	return s.r.Select(id)
}

func (s PersonService) ReadAll() ([]model.Person, error) {
	return s.r.SelectAll()
}

func (s PersonService) Update(id int, newPerson model.Person) (model.Person, error) {
	return s.r.Update(id, newPerson)
}

func (s PersonService) Delete(id int) error {
	return s.r.Delete(id)
}
