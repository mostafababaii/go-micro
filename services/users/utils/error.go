package utils

import (
	"fmt"
)

type RestError interface {
	ErrorMessage() string
	Status() int
	Message() string
	Cause() string
}

type restError struct {
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
	ErrCause   string `json:"cause"`
}

func NewRestError(status int, message string, cause string) *restError {
	return &restError{
		status,
		message,
		cause,
	}
}

func (e *restError) ErrorMessage() string {
	return fmt.Sprintf("message: %s, status: %d, cause: %s", e.ErrMessage, e.ErrStatus, e.ErrCause)
}

func (e *restError) Status() int {
	return e.ErrStatus
}

func (e *restError) Message() string {
	return e.ErrMessage
}

func (e *restError) Cause() string {
	return e.ErrCause
}
