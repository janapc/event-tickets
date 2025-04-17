package api

import (
	"encoding/json"
	"log"
	"net/http"
)

var MapErrors = map[string]int{
	"the name field is mandatory":                                            http.StatusBadRequest,
	"the description field is mandatory":                                     http.StatusBadRequest,
	"the price field cannot be less than or equal to zero":                   http.StatusBadRequest,
	"the field event_date is mandatory and should is this format DD/MM/YYYY": http.StatusBadRequest,
	"the event_date field cannot be less than current date":                  http.StatusBadRequest,
	"the currency field is mandatory":                                        http.StatusBadRequest,
	"event is not found":                                                     http.StatusNotFound,
	"the authorization is mandatory":                                         http.StatusBadRequest,
	"unauthorized user":                                                      http.StatusUnauthorized,
	"invalid token":                                                          http.StatusBadRequest,
	"you don't have permission to access this resource":                      http.StatusForbidden,
}

type OutputError struct {
	Message string `json:"message"`
}

func HandlerErrors(err error) ([]byte, int) {
	e := err.Error()
	log.Println(e)
	message := e
	statusCode := MapErrors[e]
	if MapErrors[e] == 0 {
		message = "internal server error"
		statusCode = 500
	}
	errorFormatted, _ := json.Marshal(OutputError{Message: message})
	return []byte(errorFormatted), statusCode
}
