package service

import (
	"database/sql"
)

type exampleService struct {
	db *sql.DB
}

type ExampleService interface {
	ExampleService(id string) (string, error)
	// Add function in here
}

func NewExampleService(db *sql.DB) ExampleService {
	return &exampleService{db}
}

// Write code in here
func (q *exampleService) ExampleService(id string) (string, error) {
	return "success", nil
}