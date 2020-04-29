package KeiUserUtil

import (
	"errors"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

func ValidUser(userID int) (err error){
	var user 		structs.User
	uErr := application.Db.Find(&user, userID).Error
	if uErr != nil {
		err = errors.New("user not found")
	}
	return
}