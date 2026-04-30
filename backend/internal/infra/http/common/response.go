package common

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/user/gapsi_orders_api/internal/domain"
)

type ErrorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Endpoint string `json:"endpoint"`
}

func RespondWithError(w http.ResponseWriter, r *http.Request, err error) {
	code := http.StatusInternalServerError
	message := "error interno del servidor"

	if errors.Is(err, domain.ErrUserAlreadyExists) {
		code = http.StatusBadRequest
		message = err.Error()
	} else if errors.Is(err, domain.ErrInvalidCredentials) || errors.Is(err, domain.ErrInvalidToken) || errors.Is(err, domain.ErrUserNotFound) {
		code = http.StatusUnauthorized
		message = err.Error()
	}

	RespondWithCustomError(w, r, code, message)
}

func RespondWithCustomError(w http.ResponseWriter, r *http.Request, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Code:     code,
		Message:  message,
		Endpoint: r.URL.Path,
	})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
