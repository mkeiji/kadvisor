package validators

import (
	"errors"
	util "kadvisor/server/libs/ValidationHelper"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	tagValidator    *validator.Validate
	loginRepository repository.LoginRepository
}

func (u UserValidator) Validate(obj interface{}) []error {
	errList := []error{}
	user, _ := obj.(structs.User)
	u.validateProperties(user, &errList)
	u.validateLogin(user, &errList)
	return errList
}

func (u UserValidator) validateProperties(
	user structs.User,
	errList *[]error,
) {
	u.tagValidator = validator.New()

	err := u.tagValidator.Struct(user)
	if err != nil {
		*errList = append(*errList, err)
	}
}

func (u UserValidator) validateLogin(
	user structs.User,
	errList *[]error,
) {
	_, err := u.loginRepository.FindOneByEmail(user.Login.Email)
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
