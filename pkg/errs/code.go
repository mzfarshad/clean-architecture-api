package errs

import "errors"

var (
	internalServerError = errors.New("internal server error")
	unauthorizedError   = errors.New("unauthorized error")
	forbiddenError      = errors.New("forbidden error")
	validationError     = errors.New("validation error")
	notFoundError       = errors.New("not found error")
	duplicationError    = errors.New("duplication error")
	unprocessableError  = errors.New("unprocessable error")
)

type Code interface {
	Index() int
	Err() error
}

const (
	Internal code = iota
	Unauthorized
	Forbidden
	Validation
	NotFound
	Duplication
	Unprocessable
)

type code uint8

func (c code) Index() int {
	return int(c)
}

func (c code) Err() error {
	switch c {
	case Unauthorized: // 1
		return unauthorizedError
	case Forbidden: // 2
		return forbiddenError
	case Validation: // 3
		return validationError
	case NotFound: // 4
		return notFoundError
	case Duplication: // 5
		return duplicationError
	case Unprocessable: // 6
		return unprocessableError
	default: // 0
		return internalServerError
	}
}
