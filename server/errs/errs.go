package errs

import (
	"net/http"
)

// this file use for status code task
// create new error 
type AppError struct{
	Code int
	Message string
}
// error has interface 
//type error interface {
//	Error() string
//}
func (e AppError) Error() string{
	return e.Message
}

func NewNotfoundError(message string) error {
	return AppError{
		Code: http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError() error{
	return AppError{
		Code: http.StatusInternalServerError,
		Message: "unexcepted error",
	}
}

func NewVaildationError(message string) error{
	return AppError{
		Code: http.StatusUnprocessableEntity,
		Message: message,
	}
}