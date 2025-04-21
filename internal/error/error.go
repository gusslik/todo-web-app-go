package custom_error

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	Msg        string `json:"msg"`
}

func (e ApiError) Error() string {
	return e.Msg
}
func NewApiError(statusCode int, msg string) error {
	return &ApiError{statusCode, msg}
}

func ErrorHandler(w http.ResponseWriter, err error) {
	var apiError *ApiError

	if errors.As(err, &apiError) {
		w.WriteHeader(apiError.StatusCode)
		json.NewEncoder(w).Encode(apiError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ApiError{
		StatusCode: http.StatusInternalServerError,
		Msg:        err.Error(),
	})
}
