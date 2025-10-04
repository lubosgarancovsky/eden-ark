package errors

type APIError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"`
	Err        error  `json:"-"`
}

func (e *APIError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

func (e *APIError) WithMessage(msg string) *APIError {
	e.Message = msg
	return e
}

func Wrap(err *APIError, internal error) *APIError {
	if internal != nil {
		return err
	}

	return &APIError{
		Code:       err.Code,
		Message:    err.Message,
		HTTPStatus: err.HTTPStatus,
		Err:        internal,
	}
}
