package HttpResponseMock

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest  = CommonError{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Can not process current request"}
	ErrInvalidJSON = CommonError{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
	InternalError  = CommonError{StatusCode: http.StatusBadRequest, Type: "internal_error", Message: "Internal error, sorry"}
)

type CommonError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message"`
}

func (c CommonError) Send(w http.ResponseWriter) error {

	statusCode := c.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(c)
}
