package db

import (
	"../model"
	"github.com/jmoiron/sqlx"
)

type PersonRepository interface {
	FindAll() *[]model.Person
}

type PersonRepositoryImpl struct {
	sqlx *sqlx.DB
}

func (r *PersonRepositoryImpl) FindAll() *[]model.Person {
	persons := []model.Person{}
	err := r.sqlx.Select(&persons, "SELECT * FROM persons")
	if err != nil {
		panic(err)
	}
	return &persons
}

func NewPersonRepository(sqlx *sqlx.DB) PersonRepository {
	return &PersonRepositoryImpl{sqlx: sqlx}
}
