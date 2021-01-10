package validators

import (
	"errors"
	"github.com/go-playground/validator/v10"
	util "kadvisor/server/libs/ValidationHelper"
	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
)

type EntryValidator struct {
	TagValidator    i.TagValidator
	ClassRepository i.ClassRepository
	CodeRepository  i.CodeCodeTextRepository
}

func NewEntryValidator() EntryValidator {
	return EntryValidator{
		TagValidator:    TagValidator{},
		ClassRepository: r.NewClassRepository(),
		CodeRepository:  r.NewCodeCodeTextRepository(),
	}
}

func (e EntryValidator) Validate(obj interface{}) []error {
	errList := []error{}
	entry, _ := obj.(structs.Entry)
	e.validateProperties(entry, &errList)
	e.validateClass(entry.ClassID, &errList)
	e.validateCodeType(entry.EntryTypeCodeID, &errList)
	return errList
}

func (e EntryValidator) validateProperties(
	entry structs.Entry,
	errList *[]error,
) {
	e.TagValidator.RegisterTag("ispositive", e.customIsPositive)

	err := e.TagValidator.ValidateStruct(entry)
	if err != nil {
		*errList = append(*errList, err)
	}
}

func (e EntryValidator) customIsPositive(fieldLevel validator.FieldLevel) bool {
	return fieldLevel.Field().Int() >= 0
}

func (e EntryValidator) validateClass(
	classID int,
	errList *[]error,
) {
	_, cErr := e.ClassRepository.FindOne(classID)
	if cErr != nil {
		*errList = append(
			*errList,
			errors.New(util.GetValidationMsg(
				"Entry.ClassID",
				"invalid classID",
			)),
		)
	}
}

func (e EntryValidator) validateCodeType(
	codeTypeID string,
	errList *[]error,
) {
	_, lErr := e.CodeRepository.FindOne(codeTypeID)
	if lErr != nil {
		*errList = append(
			*errList,
			errors.New(util.GetValidationMsg(
				"Entry.EntryTypeCodeID",
				"invalid lookup",
			)),
		)
	}
}
