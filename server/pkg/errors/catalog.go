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

	ErrInvalidCredentials = &APIError{
		Code:       "INVALID CREDENTIALS",
		Message:    "Invalid credentials",
		HTTPStatus: http.StatusUnauthorized,
	}

	ErrUnauthorized = &APIError{
		Code:       "UNAUTHORIZED",
		Message:    "Unauthorized",
		HTTPStatus: http.StatusUnauthorized,
	}

	ErrInvalidGrant = &APIError{
		Code:       "INVALID GRANT",
		Message:    "Invalid grant",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidScope = &APIError{
		Code:       "INVALID SCOPE",
		Message:    "Invalid scope",
		HTTPStatus: http.StatusBadRequest,
	}
)
