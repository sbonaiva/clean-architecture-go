package dto

import (
	"net/http"

	"github.com/sbonaiva/clean-architecture-go/core/domain"
)

type ErrorResponse struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
}

func ToErrorResponse(d *domain.CoreError) *ErrorResponse {

	return &ErrorResponse{
		Status:  d.Status,
		Message: d.Message,
		Fields:  d.Fields,
	}
}

func NewDefaultErrorResponse() *ErrorResponse {

	return &ErrorResponse{
		Status:  http.StatusInternalServerError,
		Message: "unexpected error",
		Fields:  nil,
	}
}
