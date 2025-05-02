package errs_test

import (
	"errors"
	"fmt"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"testing"
)

func TestNew(t *testing.T) {
	cause1 := errors.New("cause1")
	cause2 := fmt.Errorf("cause2 caused by %w", cause1)

	notFound := errs.New(errs.NotFound, "something not found")
	assert.True(t, errors.Is(notFound, errs.NotFound.Err()))
	assert.False(t, errors.Is(notFound, cause1))
	assert.False(t, errors.Is(notFound, cause2))

	notFoundCausedBy := errs.New(errs.NotFound, "something not found").CausedBy(cause2)
	assert.True(t, errors.Is(notFoundCausedBy, cause1))
	assert.True(t, errors.Is(notFoundCausedBy, cause2))
	assert.True(t, errors.Is(notFoundCausedBy, errs.NotFound.Err()))

	assert.False(t, errors.Is(notFound, notFoundCausedBy))

	var x errs.Error
	assert.True(t, errors.As(notFound, &x))
	assert.True(t, errors.As(notFoundCausedBy, &x))

}

func pathErrFirstHandler(path any) errs.Handler {
	return func(err error) errs.Error {
		if errors.Is(err, fs.ErrNotExist) {
			return errs.New(errs.NotFound, fmt.Sprintf("%v not found", path))
		} else if errors.Is(err, fs.ErrExist) {
			return errs.New(errs.Duplication, fmt.Sprintf("%v already exists", path))
		}
		return nil // Allowing next handler to be executed
	}
}
func pathErrFinisherHandler(path any) errs.Handler {
	return func(err error) errs.Error {
		if errors.Is(err, fs.ErrInvalid) {
			return errs.New(errs.Validation, fmt.Sprintf("%v is not valid", path))
		} else if errors.Is(err, fs.ErrPermission) {
			return errs.New(errs.Forbidden, fmt.Sprintf("%v is forbidden", path))
		}
		return errs.New(errs.Internal, "something went wrong") // This is a finisher
	}
}

func TestHandle(t *testing.T) {
	pathHandlers := []errs.Handler{
		pathErrFirstHandler("path"),
		pathErrFinisherHandler("path"),
	}
	type expected struct {
		code errs.Code
		msg  string
	}
	testCases := []struct {
		name     string
		input    error
		handlers []errs.Handler
		expected *expected
	}{
		{"with nil input error", nil, nil, nil},
		{"with custom input error", errs.New(errs.Forbidden, "msg"), nil, &expected{errs.Forbidden, "msg"}},
		{"with nil handler", errors.New("any error"), nil, nil},
		{"without finisher", errors.New("any error"), []errs.Handler{pathErrFirstHandler("path")}, nil},
		{"path not exists", fs.ErrNotExist, pathHandlers, &expected{errs.NotFound, "path not found"}},
		{"path exists", fs.ErrExist, pathHandlers, &expected{errs.Duplication, "path already exists"}},
		{"invalid path", fs.ErrInvalid, pathHandlers, &expected{errs.Validation, "path is not valid"}},
		{"forbidden path", fs.ErrPermission, pathHandlers, &expected{errs.Forbidden, "path is forbidden"}},
		{"path closed", fs.ErrClosed, pathHandlers, &expected{errs.Internal, "something went wrong"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := errs.Handle(tc.input, tc.handlers...)
			// Ensure the errs.Handle function
			//1. will return nil in case of nil input error,
			//2. will return exactly the same error in case of any custom input,
			//3. will wrap any other error in the result error
			assert.True(t, errors.Is(err, tc.input))
			// Check the result error type
			var got errs.Error
			if errors.As(err, &got) {
				assert.Equal(t, tc.expected.code, got.Code())
				assert.Equal(t, tc.expected.msg, got.Error())
				assert.True(t, errors.Is(err, got.Code().Err()))
			}
		})
	}
}

func TestTrace(t *testing.T) {
	notFound := errs.New(errs.NotFound, "something not found")
	assert.Equal(t, []error{errs.NotFound.Err()}, notFound.Trace())

	cause1 := errors.New("cause1")
	cause2 := fmt.Errorf("cause2 caused by (%w)", cause1)
	notFoundCausedBy := errs.New(errs.NotFound, "something not found").CausedBy(cause2)
	assert.Equal(t, []error{errs.NotFound.Err(), cause2, cause1}, notFoundCausedBy.Trace())
}
