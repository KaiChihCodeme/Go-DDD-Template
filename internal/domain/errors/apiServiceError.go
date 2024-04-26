package errors

import "net/http"

// use when API not respond 200
type ApiServerError struct {
	response      http.Response
	errStatusCode uint
	source        string
}

func (e *ApiServerError) Error() http.Response {
	return e.response
}

func NewApiServerError(response http.Response, source string) *ApiServerError {
	return &ApiServerError{
		response:      response,
		errStatusCode: uint(response.StatusCode),
		source:        source,
	}
}
