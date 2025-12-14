package repository

import (
	"database/sql"
)

type exampleRepository struct {
	db *sql.DB
}

type ExampleRepository interface {
	ExampleRepo(id string) error
	// Add function in here
}

func NewExampleRepository(db *sql.DB) ExampleRepository {
	return &exampleRepository{db}
}

// Write code in here
func (q *exampleRepository) ExampleRepo(id string) error {

	if _, err := q.db.Exec("INSERT INTO table VALUES($1)", id); err != nil {
		return err
	}

	return nil
}