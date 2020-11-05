package validators

import (
	"github.com/go-playground/validator/v10"
	app "kadvisor/server/resources/application"
)

type TagValidator struct{}

func (t TagValidator) ValidateStruct(obj interface{}) error {
	return app.Validator.Struct(obj)
}

func (t TagValidator) RegisterTag(
	tag string,
	fn validator.Func,
	callValidationEvenIfNull ...bool,
) error {
	return app.Validator.RegisterValidation(tag, fn, callValidationEvenIfNull...)
}
