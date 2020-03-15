package errors

import "fmt"

type DDDGuysError struct {
	message string
	code    DDDGuysErrorCode
}

func New(message string, code DDDGuysErrorCode) *DDDGuysError {
	return &DDDGuysError{message, code}
}

func (e DDDGuysError) Error() string {
	return fmt.Sprintf(`%s with error code: %d`, e.message, e.code)
}

func (e DDDGuysError) Code() DDDGuysErrorCode {
	return e.code
}

func Code(err error) DDDGuysErrorCode {
	if err == nil {
		return NoError
	}
	if customError, ok := err.(interface {
		Code() DDDGuysErrorCode
	}); ok {
		return customError.Code()
	}
	return Unknown
}
