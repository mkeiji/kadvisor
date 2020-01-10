package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type UserRepository struct {}

func (t *UserRepository) FindAll() ([]structs.User, error) {
	var users []structs.User
	err := application.Db.Preload("Login").Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (t *UserRepository) FindOne(id int) (structs.User, error) {
	var user structs.User
	err := application.Db.Preload("Login").Find(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (t *UserRepository) Create(user structs.User) (structs.User, error) {
	err := application.Db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
