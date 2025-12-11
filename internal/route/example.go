package route

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
}