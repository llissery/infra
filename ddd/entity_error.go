package ddd

import "fmt"

type EntityError interface {
	Code() int
	Message() string
	Error() string

	WithError(err error) EntityError
	Unwrap() error
	Is(error) bool
}

func NewEntityError(code int, msg string) EntityError {
	return &entityError{code: code, msg: msg}
}

type entityError struct {
	code int
	msg  string
	err  error
}

func (e *entityError) Code() int {
	return e.code
}

func (e *entityError) Message() string {
	return e.msg
}

func (e *entityError) Error() string {
	if e.err == nil {
		return fmt.Sprintf("EntityError: code:%v msg:%v", e.code, e.msg)
	}
	return fmt.Sprintf("EntityError: code:%v msg:%v err: %v", e.code, e.msg, e.err)
}

func (e *entityError) WithError(err error) EntityError {
	return &entityError{code: e.code, msg: e.msg, err: err}
}

func (e *entityError) Unwrap() error {
	return e.err
}

func (e *entityError) Is(err error) bool {
	if t, ok := err.(*entityError); ok {
		return e.code == t.code
	}
	return false
}
