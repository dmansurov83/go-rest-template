package db

import (
	"../model"
	"github.com/jmoiron/sqlx"
)

type PersonRepository interface {
	FindAll() (*[]model.Person, error)
	Insert(person model.Person) (*model.Person, error)
	Get(id int) (*model.Person, error)
	Delete(id int) error
}

type PersonRepositoryImpl struct {
	sqlx *sqlx.DB
}

func (r *PersonRepositoryImpl) FindAll() (*[]model.Person, error) {
	persons := []model.Person{}
	err := r.sqlx.Select(&persons, "SELECT * FROM persons")
	if err != nil {
		return nil, err
	}
	return &persons, nil
}

func (r *PersonRepositoryImpl) Get(id int) (*model.Person, error) {
	person := model.Person{}
	err := r.sqlx.Get(&person, "SELECT * FROM persons WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepositoryImpl) Insert(person model.Person) (*model.Person, error) {
	rows, err := r.sqlx.NamedQuery("INSERT INTO persons (name, age) VALUES (:name, :age)", person)
	if err != nil {
		return nil, err
	}
	var id int
	rows.Next()
	err = rows.Scan(&id)
	if err != nil {
		return nil, err
	}
	return r.Get(id)
}

func (r *PersonRepositoryImpl) Delete(id int) error {
	_, err := r.sqlx.Exec("DELETE FROM persons WHERE id=$1", id)
	return err
}

func NewPersonRepository(sqlx *sqlx.DB) PersonRepository {
	return &PersonRepositoryImpl{sqlx: sqlx}
}
