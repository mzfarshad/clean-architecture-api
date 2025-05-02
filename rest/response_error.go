package rest

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"net/http"
)

type responseErr struct {
	Code     errs.Code `json:"code"`
	Type     string    `json:"type"`
	Messages []string  `json:"messages"`
	Debug    []string  `json:"debug,omitempty"`
}

func (e *responseErr) from(err errs.Error) *responseErr {
	if e == nil {
		return new(responseErr).from(err)
	}
	e.Code = err.Code()
	e.Type = err.Code().Err().Error()
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, fieldError := range validationErrs {
			e.Messages = append(e.Messages, fieldError.Error())
		}
	} else {
		e.Messages = append(e.Messages, err.Error())
	}
	// TODO: add app config
	//if config.Get().App.Debug {
	//	trace := err.Trace()
	//	for i := range trace {
	//		e.Debug = append(e.Debug, trace[i].Error())
	//	}
	//}
	return e
}

func (e *responseErr) httpStatus() int {
	if e == nil {
		return http.StatusOK
	}
	switch e.Code {
	case errs.Unauthorized:
		return http.StatusUnauthorized
	case errs.Forbidden:
		return http.StatusForbidden
	case errs.Validation:
		return http.StatusBadRequest
	case errs.NotFound:
		return http.StatusNotFound
	case errs.Duplication:
		return http.StatusConflict
	case errs.Unprocessable:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
