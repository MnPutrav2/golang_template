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