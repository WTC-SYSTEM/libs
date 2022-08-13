package apperror

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNotFound = NewAppError("not found", WTC000003, "")
	EmailIsRegistered = NewAppError("email is already registered", WTC000004, "")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func NewAppError(message string, code WtcError, developerMessage string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Code:             code.String(),
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}
	return bytes
}

func BadRequestError(message string) *AppError {
	return NewAppError(message, WTC000002, "Some thing wrong with user data")
}

func SystemError(developerMessage string) *AppError {
	return NewAppError("System error", WTC000001, developerMessage)
}

func FromError(err error) *AppError {
	return NewAppError("Something went wrong...", WTC000001, err.Error())
}

func APIError(code WtcError, message, developerMessage string) *AppError {
	return NewAppError(message, code, developerMessage)
}
