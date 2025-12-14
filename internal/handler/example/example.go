package handler

import (
	"clean-arsitektur/internal/model"
	logging "clean-arsitektur/pkg/logging"
	"clean-arsitektur/pkg/middleware"
	"database/sql"
	"encoding/json"
	"net/http"
)

func ExampleController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "GET":
			// Write method get
			// Example
			handlerName(db)(w, r)
		case "POST":
			// Write method post
		case "PUT":
			// Write method put
		case "DELETE":
			// Write method delete
		}

	}
}

func handlerName(_ *sql.DB) http.HandlerFunc {
	return middleware.CORS(
		middleware.RateLimiter(1, 1, func(w http.ResponseWriter, r *http.Request) {
			// Write code in here
			res, _ := json.Marshal(model.ResponseMessage{Status: "success", Message: "message"})
			logging.Log("message", "INFO", r)
			w.WriteHeader(200)
			w.Write(res)
		}),
	)
}
