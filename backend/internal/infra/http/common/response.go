package common

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Endpoint string `json:"endpoint"`
}

func RespondWithError(w http.ResponseWriter, r *http.Request, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Code:     code,
		Message:  message,
		Endpoint: r.URL.Path,
	})
}
