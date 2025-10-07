package errors

import (
	"net/http"

	"github.com/lubosgarancovsky/go-kit/api_err"
)

var (
	ErrParameterMissing = &api_err.ApiError{
		Code:       "PARAMETER MISSING",
		Message:    "Parameter missing",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidGrant = &api_err.ApiError{
		Code:       "INVALID GRANT",
		Message:    "Invalid grant",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidScope = &api_err.ApiError{
		Code:       "INVALID SCOPE",
		Message:    "Invalid scope",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidCredentials = &api_err.ApiError{
		Code:       "INVALID CREDENTIALS",
		Message:    "Entered credentials are invalid",
		HTTPStatus: http.StatusUnauthorized,
	}
)
