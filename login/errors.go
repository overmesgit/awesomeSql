package login

import "fmt"

type Error struct {
	orig error
	msg  string
	code ErrorCode
}

func (e *Error) Error() string {
	if e.orig != nil {
		return fmt.Sprintf("%s: %v", e.msg, e.orig)
	}

	return e.msg
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func WrapError(orig error, msg string, code ErrorCode) *Error {
	if code > InternalError {
		code = InternalError
	}
	return &Error{orig, msg, code}
}

type ErrorCode uint

const (
	UserNotFoundError ErrorCode = iota
	UserAlreadyExistError
	ValidationError

	// InternalError should be last for boundary check
	InternalError
)
