package validators

import (
	"errors"
	"kadvisor/server/libs/KeiPassUtil"
	util "kadvisor/server/libs/ValidationHelper"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"

	"github.com/go-playground/validator/v10"
)

type LoginValidator struct {
	tagValidator    *validator.Validate
	loginRepository repository.LoginRepository
}

func (l LoginValidator) Validate(obj interface{}) []error {
	errList := []error{}
	login, _ := obj.(structs.Login)
	l.validateLoginExists(login, &errList)
	l.validatePassword(login, &errList)
	return errList
}

func (l LoginValidator) validateLoginExists(
	login structs.Login,
	errList *[]error,
) {
	_, err := l.loginRepository.FindOneByEmail(login.Email)
	if err != nil {
		*errList = append(
			*errList,
			errors.New(util.GetValidationMsg(
				"Login.Email",
				"invalid email",
			)),
		)
	}
}

func (l LoginValidator) validatePassword(
	login structs.Login,
	errList *[]error,
) {
	stored, _ := l.loginRepository.FindOneByEmail(login.Email)
	if !KeiPassUtil.IsValidPassword(stored.Password, login.Password) {
		*errList = append(
			*errList,
			errors.New(util.GetValidationMsg(
				"Login.Password",
				"wrong password",
			)),
		)
	}
}
