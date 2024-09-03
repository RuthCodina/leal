package helpers

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest     = handlerError{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Cannot process current request"}
	ErrInvalidJSON    = handlerError{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
	ErrReadingRequest = handlerError{StatusCode: http.StatusInternalServerError, Type: "read_request", Message: "error trying to read the request"}
	ErrParsingDate    = handlerError{StatusCode: http.StatusInternalServerError, Type: "parser_date", Message: "error trying to parse payment date"}
)

type handlerError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (e handlerError) Send(w http.ResponseWriter) error {
	statusCode := e.StatusCode
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(e)
}
