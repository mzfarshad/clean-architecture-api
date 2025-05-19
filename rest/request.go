package rest

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/mzfarshad/music_store_api/pkg/errs"
)

var (
	validate         = validator.New()
	errParsingBody   = errors.New("failed to parse request body")
	errParsingParams = errors.New("failed to parse request params")
)

type Request[T idto] struct{}

// Parse binds the request body of *fiber.Ctx into the given rest.DTO struct and validates it.
// It returns a non-nil pointer of the given input or error.
func (Request[T]) Parse(ctx *fiber.Ctx) (*T, error) {
	var body *T
	if err := ctx.BodyParser(&body); err != nil {
		return nil, errs.New(errs.Unprocessable, "unprocessable request body").CausedBy(err)
	}
	if body == nil {
		return nil, errParsingBody
	}
	if err := validate.Struct(body); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}
		errorList := err.(validator.ValidationErrors)
		return nil, errs.New(errs.Validation, errorList.Error()).CausedBy(err)
	}
	return body, nil
}

// ParseParams binds the request params of *fiber.Ctx into the given DTO struct and validates it.
// It returns a non-nil pointer of the given params or error.
func (Request[T]) ParseParams(ctx *fiber.Ctx) (*T, error) {
	var params T
	if err := ctx.ParamsParser(&params); err != nil {
		return nil, errs.New(errs.Unprocessable, "unprocessable request params").CausedBy(err)
	}
	if &params == nil {
		return nil, errParsingParams
	}
	if err := validate.Struct(params); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}
		errorList := err.(validator.ValidationErrors)
		return nil, errs.New(errs.Validation, errorList.Error()).CausedBy(err)
	}
	return &params, nil
}

// ParseQueries binds the request Queries string of *fiber.Ctx into the given DTO struct and validates it.
// It returns a non-nil pointer of the given queries or error.
func (Request[T]) ParseQueries(ctx *fiber.Ctx) (*T, error) {
	var queries T
	if err := ctx.QueryParser(&queries); err != nil {
		return nil, errs.New(errs.Unprocessable, "unprocessable request params").CausedBy(err)
	}
	if &queries == nil {
		return nil, errParsingParams
	}
	if err := validate.Struct(queries); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil, err
		}
		errorList := err.(validator.ValidationErrors)
		return nil, errs.New(errs.Validation, errorList.Error()).CausedBy(err)
	}
	return &queries, nil
}
