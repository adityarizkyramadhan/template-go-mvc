package utils

import (
	"fmt"
	"strings"
)

type CustomError struct {
	Key     string
	Message string
}

type ParsingError struct {
	Key        string `json:"key"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

var (
	ErrInternalServer = "internal_server_error"
	ErrNotFound       = "not_found"
	ErrConflict       = "conflict"
	ErrBadRequest     = "bad_request"
	ErrUnauthorized   = "unauthorized"
	ErrForbidden      = "forbidden"
	ErrValidation     = "validation"
	ErrUnknown        = "unknown"
)

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Key, e.Message)
}

func NewError(key, message string) error {
	return CustomError{
		Key:     key,
		Message: message,
	}
}

func ParseError(error string) ParsingError {
	var e CustomError
	splitError := strings.Split(error, ": ")
	if len(splitError) == 2 {
		e.Key = splitError[0]
		e.Message = splitError[1]
	} else {
		e.Key = ErrUnknown
		e.Message = error
	}
	var parseError ParsingError
	parseError.Key = e.Key
	parseError.Message = e.Message

	switch e.Key {
	case ErrInternalServer:
		parseError.StatusCode = 500
	case ErrNotFound:
		parseError.StatusCode = 404
	case ErrConflict:
		parseError.StatusCode = 409
	case ErrBadRequest:
		parseError.StatusCode = 400
	case ErrUnauthorized:
		parseError.StatusCode = 401
	case ErrForbidden:
		parseError.StatusCode = 403
	case ErrValidation:
		parseError.StatusCode = 422
	case ErrUnknown:
		parseError.StatusCode = 500
	default:
		parseError.StatusCode = 500
	}
	return parseError
}
