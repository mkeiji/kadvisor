package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type LoginRepository struct {}

func (l *LoginRepository) FindOneByEmail(email string) structs.Login {
	var login structs.Login
	application.Db.Where("email=?", email).First(&login)
	return login
}

func (l *LoginRepository) UpdateLoginStatus(login structs.Login, isLoggedIn bool) structs.Login {
	var storedLogin structs.Login
	application.Db.Find(&storedLogin, login.ID).Update("IsLoggedIn", isLoggedIn)
	return storedLogin
}