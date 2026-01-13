package errors

import (
	"net/http"

	"github.com/lubosgarancovsky/go-kit"
)

var (
	ErrParameterMissing = &go_kit.ApiError{
		Code:       "PARAMETER MISSING",
		Message:    "Parameter missing",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidGrant = &go_kit.ApiError{
		Code:       "INVALID GRANT",
		Message:    "Invalid grant",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidScope = &go_kit.ApiError{
		Code:       "INVALID SCOPE",
		Message:    "Invalid scope",
		HTTPStatus: http.StatusBadRequest,
	}

	ErrInvalidCredentials = &go_kit.ApiError{
		Code:       "INVALID CREDENTIALS",
		Message:    "Entered credentials are invalid",
		HTTPStatus: http.StatusUnauthorized,
	}
)
