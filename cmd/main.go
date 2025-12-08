package main

import (
	"clean-arsitektur/internal/config"
	"net/http"
	"os"
)

func main() {
	db := config.Database()
	defer db.Close()

	http.ListenAndServe(os.Getenv("APP_LISTEN"), nil)
}
