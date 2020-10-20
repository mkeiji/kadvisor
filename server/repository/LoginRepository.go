package repository

import (
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
)

type LoginRepository struct{}

func (l *LoginRepository) FindOneByEmail(email string) (s.Login, error) {
	var login s.Login

	err := app.Db.Where("email=?", email).First(&login).Error
	if err != nil {
		return login, err
	}
	return login, nil
}

func (l *LoginRepository) Update(
	login s.Login,
) (s.Login, error) {
	stored, err := l.findOne(login.ID)
	if err == nil {
		err = app.Db.Model(&stored).Updates(login).Error
	}

	return stored, err
}

func (l *LoginRepository) UpdateLoginStatus(
	login s.Login,
	isLoggedIn bool,
) (s.Login, error) {
	var storedLogin s.Login

	err := app.Db.Find(&storedLogin, login.ID).Update("IsLoggedIn", isLoggedIn).Error
	if err != nil {
		return storedLogin, err
	}

	return storedLogin, nil
}

func (l *LoginRepository) findOne(
	id int,
) (s.Login, error) {
	var storedLogin s.Login
	err := app.Db.Where("id=?", id).First(&storedLogin).Error
	return storedLogin, err
}
