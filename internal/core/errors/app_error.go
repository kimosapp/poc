package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Message     string `json:"message"`
	Code        string `json:"code"`
	Description string `json:"description"`
	HTTPStatus  int
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NotFoundError represents an error when a resource is not found.
type NotFoundError struct {
	*AppError
}

func NewNotFoundError(message, description string, code ErrorCode) *NotFoundError {
	return &NotFoundError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusNotFound,
		},
	}
}

type UnauthorizedError struct {
	*AppError
}

func NewUnauthorizedError(message, description string, code ErrorCode) *UnauthorizedError {
	return &UnauthorizedError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusUnauthorized,
		},
	}
}

type BadRequestError struct {
	*AppError
}

func NewBadRequestError(message, description string, code ErrorCode) *BadRequestError {
	return &BadRequestError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusBadRequest,
		},
	}
}

type InternalServerError struct {
	*AppError
}

func NewInternalServerError(message, description string, code ErrorCode) *InternalServerError {
	return &InternalServerError{
		AppError: &AppError{
			Message:     message,
			Code:        string(code),
			Description: description,
			HTTPStatus:  http.StatusInternalServerError,
		},
	}
}
