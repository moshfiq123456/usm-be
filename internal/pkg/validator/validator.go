package validator

import (
	"github.com/go-playground/validator/v10"
	"gofr.dev/pkg/gofr"
)

var validate = validator.New()

// BindAndValidate binds request body and validates struct tags
func BindAndValidate(ctx *gofr.Context, req interface{}) error {
	if err := ctx.Bind(req); err != nil {
		return err
	}

	return validate.Struct(req)
}
