package validators

import (
	"errors"
	"kadvisor/server/libs/KeiPassUtil"
	util "kadvisor/server/libs/ValidationHelper"
	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
	s "kadvisor/server/repository/structs"
)

type LoginValidator struct {
	TagValidator    i.TagValidator
	LoginRepository i.LoginRepository
}

func NewLoginValidator() LoginValidator {
	return LoginValidator{
		TagValidator:    TagValidator{},
		LoginRepository: r.NewLoginRepository(),
	}
}

func (l LoginValidator) Validate(obj interface{}) []error {
	errList := []error{}
	login, _ := obj.(s.Login)
	l.validateLoginExists(login, &errList)
	return errList
}

func (l LoginValidator) validateLoginExists(
	login s.Login,
	errList *[]error,
) {
	stored, err := l.LoginRepository.FindOneByEmail(login.Email)
	if err != nil {
		*errList = append(
			*errList,
			errors.New(util.GetValidationMsg(
				"Login.Email",
				"invalid email",
			)),
		)
	} else {
		l.validatePassword(stored, login, errList)
	}
}

func (l LoginValidator) validatePassword(
	stored s.Login,
	claim s.Login,
	errList *[]error,
) {
	if !KeiPassUtil.IsValidPassword(stored.Password, claim.Password) {
		*errList = append(
			*errList,
			errors.New(util.GetValidationMsg(
				"Login.Password",
				"wrong password",
			)),
		)
	}
}
