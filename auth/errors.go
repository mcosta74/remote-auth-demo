package auth

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Code int
	Err  error
}

func (e *APIError) Error() string {
	return e.Err.Error()
}

func (e *APIError) StatusCode() int {
	if e.Code == 0 {
		return http.StatusInternalServerError
	}
	return e.Code
}

func (e *APIError) MarshalJSON() ([]byte, error) {
	tmp := struct {
		Error string `json:"error"`
	}{
		Error: e.Error(),
	}
	return json.Marshal(tmp)
}
