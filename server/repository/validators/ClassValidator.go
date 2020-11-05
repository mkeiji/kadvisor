package validators

import (
	i "kadvisor/server/repository/interfaces"
	s "kadvisor/server/repository/structs"
)

type ClassValidator struct {
	TagValidator i.TagValidator
}

func NewClassValidator() ClassValidator {
	return ClassValidator{
		TagValidator: TagValidator{},
	}
}

func (c ClassValidator) Validate(obj interface{}) []error {
	errList := []error{}
	class, _ := obj.(s.Class)
	c.validateProperties(class, &errList)
	return errList
}

func (c ClassValidator) validateProperties(
	class s.Class,
	errList *[]error,
) {
	err := c.TagValidator.ValidateStruct(class)
	if err != nil {
		*errList = append(*errList, err)
	}
}
