package services

import "../db"
import "../model"

type PersonService interface {
	FindAll() (*[]model.Person, error)
	Insert(person model.Person) (*model.Person, error)
	Get(id int) (*model.Person, error)
	Delete(id int) error
}

type PersonServiceImpl struct {
	repository db.PersonRepository
}

func (service *PersonServiceImpl) FindAll() (*[]model.Person, error) {
	return service.repository.FindAll()
}

func (service *PersonServiceImpl) Insert(person model.Person) (*model.Person, error) {
	return service.repository.Insert(person)
}

func (service *PersonServiceImpl) Get(id int) (*model.Person, error) {
	return service.repository.Get(id)
}

func (service *PersonServiceImpl) Delete(id int) error {
	return service.repository.Delete(id)
}

func NewPersonService(repository db.PersonRepository) PersonService {
	return &PersonServiceImpl{repository: repository}
}
