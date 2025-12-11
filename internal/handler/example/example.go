package handler

import (
	"clean-arsitektur/internal/model"
	logging "clean-arsitektur/pkg/logging"
	"clean-arsitektur/pkg/middleware"
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
}