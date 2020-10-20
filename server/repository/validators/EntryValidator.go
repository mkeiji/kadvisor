package validators

import (
	"errors"
	util "kadvisor/server/libs/ValidationHelper"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"

	"github.com/go-playground/validator/v10"
)

type EntryValidator struct {
	tagValidator    *validator.Validate
	classRepository repository.ClassRepository
	codeRepository  repository.CodeCodeTextRepository
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
	e.tagValidator = validator.New()
	e.tagValidator.RegisterValidation("ispositive", e.customIsPositive)

	err := e.tagValidator.Struct(entry)
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
	_, cErr := e.classRepository.FindOne(classID)
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
	_, lErr := e.codeRepository.FindOne(codeTypeID)
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
