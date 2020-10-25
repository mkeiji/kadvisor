package validators

import (
	"kadvisor/server/repository/structs"

	"github.com/go-playground/validator/v10"
)

type ClassValidator struct {
	tagValidator *validator.Validate
}

func (c ClassValidator) Validate(obj interface{}) []error {
	errList := []error{}
	class, _ := obj.(structs.Class)
	c.validateProperties(class, &errList)
	return errList
}

func (c ClassValidator) validateProperties(
	class structs.Class,
	errList *[]error,
) {
	c.tagValidator = validator.New()

	err := c.tagValidator.Struct(class)
	if err != nil {
		*errList = append(*errList, err)
	}
}
