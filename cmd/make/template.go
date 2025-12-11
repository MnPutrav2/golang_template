package make

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Template(dir, name, ty string) {
	module := moduleReader()

	hdlTemp := `package handler

import (
	"` + module + `/internal/model"
	logging "` + module + `/pkg/logging"
	"` + module + `/pkg/middleware"
	"database/sql"
	"encoding/json"
	"net/http"
)

func HandlerName(db *sql.DB) http.HandlerFunc {
	return middleware.CORS(
		middleware.RateLimiter(1, 1, func(w http.ResponseWriter, r *http.Request) {
			// Write code in here
			res, _ := json.Marshal(model.ResponseMessage{Status: "success", Message: "message"})
			logging.Log("message", "INFO", r)
			w.WriteHeader(200)
			w.Write(res)
		}),
	)
}`

	routTemp := `package route

import (
	"database/sql"
	"net/http"
)

func RouteName(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			// Write method get
		case "POST":
			// Write method post
		case "PUT":
			// Write method put
		case "DELETE":
			// Write method delete
		}

	}
}`

	srvImplTemp := `package service

import (
	"database/sql"
)

type ` + name + `Service struct {
	db *sql.DB
}

type ` + capitalize(name) + `Service interface {
	ExampleService(id string) (string, error)
	// Add function in here
}

func New` + capitalize(name) + `Service(db *sql.DB) ` + capitalize(name) + `Service {
	return &` + name + `Service{db}
}

// Write code in here`

	servTemp := `package service

func (q *` + name + `Service) ExampleService(id string) (string, error) {

	// Write code in here

	return "success", nil
}`

	repoImplTemp := `package repository

import (
	"database/sql"
)

type ` + name + `Repository struct {
	db *sql.DB
}

type ` + capitalize(name) + `Repository interface {
	ExampleRepo(id string) error
	// Add function in here
}

func New` + capitalize(name) + `Repository(db *sql.DB) ` + capitalize(name) + `Repository {
	return &` + name + `Repository{db}
}

// Write code in here`

	repoTemp := `package repository

func (q *` + name + `Repository) ExampleRepo(id string) error {

	if _, err := q.db.Exec("INSERT INTO test VALUES($1)", id); err != nil {
		return err
	}

	return nil
}`

	switch ty {
	case "handler":
		handleTemp := process(hdlTemp, "handler", dir, name)
		fmt.Println(handleTemp)
		return
	case "service":
		serviceImplTemp := process(srvImplTemp, "service", dir, name+"_impl")
		serviceTemp := process(servTemp, "service", dir, name)
		fmt.Println(serviceImplTemp)
		fmt.Println(serviceTemp)
		return
	case "route":
		routeTemp := process(routTemp, "route", "", name)
		fmt.Println(routeTemp)
		return
	case "repository":
		repositoryImplTemp := process(repoImplTemp, "repository", dir, name+"_impl")
		repositoryTemp := process(repoTemp, "repository", dir, name)
		fmt.Println(repositoryImplTemp)
		fmt.Println(repositoryTemp)
		return
	case "all":
		handleTemp := process(hdlTemp, "handler", dir, name)
		routeTemp := process(routTemp, "route", "", name)
		serviceImplTemp := process(srvImplTemp, "service", dir, name+"_impl")
		serviceTemp := process(servTemp, "service", dir, name)
		repositoryImplTemp := process(repoImplTemp, "repository", dir, name+"_impl")
		repositoryTemp := process(repoTemp, "repository", dir, name)
		fmt.Println(handleTemp)
		fmt.Println(routeTemp)
		fmt.Println(serviceImplTemp)
		fmt.Println(serviceTemp)
		fmt.Println(repositoryImplTemp)
		fmt.Println(repositoryTemp)
		return
	default:
		fmt.Println("invalid command")
		return
	}

}

func process(template, path, dir, name string) string {
	file := name + ".go"

	os.MkdirAll("internal/"+path+dir, 0o755)
	handlePath := "internal/" + path + dir
	save := filepath.Join(handlePath, file)

	os.WriteFile(save, []byte(template), 0o644)
	return "Created:" + save
}

func moduleReader() string {
	file, err := os.Open("go.mod")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "module ") {
			moduleName := strings.TrimSpace(strings.TrimPrefix(line, "module"))
			return moduleName
		}
	}

	return ""
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}
