package validators

import (
	"errors"
	util "kadvisor/server/libs/ValidationHelper"
	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
	s "kadvisor/server/repository/structs"
)

type UserValidator struct {
	TagValidator    i.TagValidator
	LoginRepository i.LoginRepository
}

func NewUserValidator() UserValidator {
	return UserValidator{
		TagValidator:    TagValidator{},
		LoginRepository: r.NewLoginRepository(),
	}
}

func (u UserValidator) Validate(obj interface{}) []error {
	errList := []error{}
	user, _ := obj.(s.User)
	u.validateProperties(user, &errList)
	u.validateLogin(user, &errList)
	return errList
}

func (u UserValidator) validateProperties(
	user s.User,
	errList *[]error,
) {
	err := u.TagValidator.ValidateStruct(user)
	if err != nil {
		*errList = append(*errList, err)
	}
}

func (u UserValidator) validateLogin(
	user s.User,
	errList *[]error,
) {
	_, err := u.LoginRepository.FindOneByEmail(user.Login.Email)
	if err == nil {
		*errList = append(
			*errList,
			errors.New(util.GetValidationMsg(
				"User.Login.Email",
				"email already exists",
			)),
		)
	}
}
