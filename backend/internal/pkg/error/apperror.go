package error

import "fmt"

type Code string

const (
	CodeNotFound   Code = "NOT_FOUND"
	CodeConflict   Code = "CONFLICT"
	CodeValidation Code = "VALIDATION"
	CodeForbidden  Code = "FORBIDDEN"
	CodeInternal   Code = "INTERNAL"
)

type AppError struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code Code, message string, underlying error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     underlying,
	}
}

func NotFound(message string, underlying error) *AppError {
	return New(CodeNotFound, message, underlying)
}

func Conflict(message string, underlying error) *AppError {
	return New(CodeConflict, message, underlying)
}

func Validation(message string, underlying error) *AppError {
	return New(CodeValidation, message, underlying)
}

func Forbidden(message string, underlying error) *AppError {
	return New(CodeForbidden, message, underlying)
}

func Internal(message string, underlying error) *AppError {
	return New(CodeInternal, message, underlying)
}
