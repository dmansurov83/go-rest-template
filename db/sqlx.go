package db

import (
	"../env"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDB(config *env.Config) *sqlx.DB {
	db, err := sqlx.Connect(config.DbDriver, config.DbSource)
	if err != nil {
		panic(err)
	}
	// todo add schema migration https://github.com/golang-migrate/migrate
	schema := `CREATE TABLE persons (
  	id INTEGER PRIMARY KEY AUTOINCREMENT,
    name text,
    age integer);`

	// execute a query on the server
	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}
	return db
}
