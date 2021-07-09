package www

import (
	"github.com/go-playground/validator"
	"go-http-boilerplate/src/domain"
)

type argumentValidator struct {
	validator *validator.Validate
}

func newArgumentValidator() *argumentValidator {
	return &argumentValidator{
		validator: validator.New(),
	}
}

func (av *argumentValidator) Validate(i interface{}) error {
	if err := av.validator.Struct(i); err != nil {
		return domain.ErrBadRequest
	}

	return nil
}
