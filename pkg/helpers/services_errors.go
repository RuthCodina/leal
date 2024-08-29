package helpers

import (
	"errors"
	"fmt"
)

const (
	ErrCodeDuplicateKey        = "duplicate_key"
	ErrCodeInternalServerError = "internal_server_error"
	ErrCodeInvalidParams       = "invalid_params"
	ErrCodeNotFound            = "not_found"
	ErrCodeTimeout             = "timeout"
)

var (
	ErrDuplicateKey = errors.New("duplicate key error")
	ErrIncorrectID  = errors.New("incorrect id error")
	ErrNotFound     = errors.New("record not found error")
	ErrTimeout      = errors.New("timeout error")
)

// Error is a custom error type that implements the error interface
type Error struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

// NewError creates a new Error with the given code and message.
func NewError(code string, msg string) Error {
	return Error{
		Code: code,
		Msg:  msg,
	}
}

// Error returns a string representation of the error. It is part of the error interface.
func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Msg)
}
