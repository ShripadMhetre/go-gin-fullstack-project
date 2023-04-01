package errors

import (
	"fmt"
)

type ErrorEntity struct {
	StatusCode int
	ErrorCode  string
	Message    string
}

func (err *ErrorEntity) Error() string {
	return fmt.Sprintf("StatusCode: %d, ErrorCode: %s, Message: %s", err.StatusCode, err.ErrorCode, err.Message)
}
