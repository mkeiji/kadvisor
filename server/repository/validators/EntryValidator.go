package validators

import (
	"errors"
	util "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"

	"github.com/go-playground/validator/v10"
)

type EntryValidator struct {
	classRepository repository.ClassRepository
	codeRepository  repository.CodeCodeTextRepository
	validator       *validator.Validate
}

var errList []error

func (e *EntryValidator) Validate(entry structs.Entry) []error {
	errList = []error{}
	e.validateProperties(entry)
	e.validateClass(entry.ClassID)
	e.validateCodeType(entry.EntryTypeCodeID)
	return errList
}

func (e *EntryValidator) validateProperties(entry structs.Entry) {
	e.validator = validator.New()
	e.validator.RegisterValidation("ispositive", e.customIsPositive)

	err := e.validator.Struct(entry)
	if err != nil {
		errList = append(errList, err)
	}
}

func (e *EntryValidator) customIsPositive(fieldLevel validator.FieldLevel) bool {
	return fieldLevel.Field().Int() >= 0
}

func (e *EntryValidator) validateClass(classID int) {
	_, cErr := e.classRepository.FindOne(classID)
	if cErr != nil {
		errList = append(errList, errors.New(util.AsValidationMsg("invalid classID")))
	}
}

func (e *EntryValidator) validateCodeType(codeTypeID string) {
	_, lErr := e.codeRepository.FindOne(codeTypeID)
	if lErr != nil {
		errList = append(
			errList,
			errors.New(util.AsValidationMsg("invalid lookup @ entryTypeCodeID")),
		)
	}
}
