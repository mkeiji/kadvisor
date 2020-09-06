package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type LoginRepository struct{}

func (l *LoginRepository) FindOneByEmail(email string) (structs.Login, error) {
	var login structs.Login

	err := application.Db.Where("email=?", email).First(&login).Error
	if err != nil {
		return login, err
	}
	return login, nil
}

func (l *LoginRepository) Update(
	login structs.Login) (structs.Login, error) {
	var storedLogin structs.Login

	err := application.Db.Find(&storedLogin, login.ID).Updates(login).Error
	return storedLogin, err
}

func (l *LoginRepository) UpdateLoginStatus(login structs.Login, isLoggedIn bool) (structs.Login, error) {
	var storedLogin structs.Login

	err := application.Db.Find(&storedLogin, login.ID).Update("IsLoggedIn", isLoggedIn).Error
	if err != nil {
		return storedLogin, err
	}

	return storedLogin, nil
}
