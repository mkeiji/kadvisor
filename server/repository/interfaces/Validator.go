package interfaces

import "github.com/go-playground/validator/v10"

//go:generate mockgen -destination=mocks/mock_validator.go -package=mocks . Validator
type Validator interface {
	Validate(obj interface{}) []error
}

//go:generate mockgen -destination=mocks/mock_tagValidator.go -package=mocks . TagValidator
type TagValidator interface {
	ValidateStruct(obj interface{}) error
	RegisterTag(tag string, fn validator.Func, callValidationEvenIfNull ...bool) error
}
