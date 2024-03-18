package api

import (
	"log/slog"
	"net/http"
)

var MapErrors = map[string]int{
	"client is not found":             http.StatusNotFound,
	"client already exists":           http.StatusConflict,
	"the name field is mandatory":     http.StatusBadRequest,
	"the email field is mandatory":    http.StatusBadRequest,
	"the event_id field is mandatory": http.StatusBadRequest,
}

type OutputError struct {
	Message string `json:"message"`
}

func HandlerErrors(err error) (OutputError, int) {
	e := err.Error()
	slog.Error("Error: " + e)
	message := e
	statusCode := MapErrors[e]
	if MapErrors[e] == 0 {
		message = "internal server error"
		statusCode = http.StatusInternalServerError
	}
	return OutputError{
		Message: message,
	}, statusCode
}
