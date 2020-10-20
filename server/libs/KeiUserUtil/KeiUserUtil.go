package KeiUserUtil

import (
	"errors"
	util "kadvisor/server/libs/ValidationHelper"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

func ValidUser(userID int) (err error) {
	var user structs.User
	uErr := application.Db.Where("id=?", userID).First(&user).Error
	if uErr != nil {
		err = errors.New(util.GetValidationMsg("User.ID", "invalid user id"))
	}
	return
}
