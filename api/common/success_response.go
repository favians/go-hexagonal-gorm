package common

import "net/http"

type SuccessResponseCode string

//List of success response status
const (
	Success SuccessResponseCode = "success"
)

//SuccessResponse default payload response
type SuccessResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
}

//NewSuccessResponse create new success payload
func NewSuccessResponse(data interface{}) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		"Success",
		data,
	}
}

//NewSuccessResponse create new success payload
func NewSuccessResponseWithoutData() (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		"Success",
		map[string]interface{}{},
	}
}
