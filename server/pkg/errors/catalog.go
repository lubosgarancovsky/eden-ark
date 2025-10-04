package errors

import "net/http"

var (
	ErrBadRequest = &APIError{
		Code:       "BAD REQUEST",
		Message:    "Bad Request",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrParameterMissing = &APIError{
		Code:       "PARAMETER MISSING",
		Message:    "Parameter missing",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidUUID = &APIError{
		Code:       "INVALID UUID",
		Message:    "Invalid UUID",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInternalServer = &APIError{
		Code:       "INTERNAL SERVER ERROR",
		Message:    "Internal server error",
		HTTPStatus: http.StatusInternalServerError,
	}

	ErrNotFound = &APIError{
		Code:       "NOT FOUND",
		Message:    "Not found",
		HTTPStatus: http.StatusNotFound,
	}
)
