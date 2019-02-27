package services

import "../db"
import "../model"

type PersonService interface {
	FindAll() *[]model.Person
}

type PersonServiceImpl struct {
	repository db.PersonRepository
}

func (service *PersonServiceImpl) FindAll() *[]model.Person {
	return service.repository.FindAll()
}

func NewPersonService(repository db.PersonRepository) PersonService {
	return &PersonServiceImpl{repository: repository}
}
