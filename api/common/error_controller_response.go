package common

import "net/http"

type errorControllerResponseCode string

const (
	ErrBadRequest errorControllerResponseCode = "bad_request"
	ErrForbidden  errorControllerResponseCode = "forbidden"
)

//ControllerResponse default payload response
type ControllerResponse struct {
	Code    errorControllerResponseCode `json:"code"`
	Message string                      `json:"message"`
	Data    interface{}                 `json:"data"`
}

//NewBadRequestResponse bad request format response
func NewBadRequestResponse() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		ErrBadRequest,
		"Bad request",
		map[string]interface{}{},
	}
}

//NewForbiddenResponse default for Forbidden error response
func NewForbiddenResponse() (int, ControllerResponse) {
	return http.StatusForbidden, ControllerResponse{
		ErrForbidden,
		"Forbidden",
		map[string]interface{}{},
	}
}
