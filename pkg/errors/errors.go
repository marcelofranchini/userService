package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func NewAppError(code int, message string) error {
	return &AppError{Code: code, Message: message}
}
