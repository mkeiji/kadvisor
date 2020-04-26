package repository

import (
	"kadvisor/server/repository/mappers"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type UserRepository struct {
	mapper mappers.UserMapper
}

func (t *UserRepository) FindAll(preloaded bool) ([]structs.User, error) {
	var users []structs.User

	if preloaded {
		err := application.Db.Preload(
			"Login").Preload(
			"Entries").Preload(
			"Classes").Find(&users).Error
		if err != nil {
			return users, err
		}
	} else {
		err := application.Db.Preload("Login").Find(&users).Error
		if err != nil {
			return users, err
		}
	}

	return t.mapper.MapSubClassesOnLoad(users, application.Db), nil
}

func (t *UserRepository) FindOne(id int, preloaded bool) (structs.User, error) {
	var user structs.User

	if preloaded {
		err := application.Db.Preload(
			"Login").Preload(
			"Entries").Preload(
			"Classes").Find(&user, id).Error
		if err != nil {
			return user, err
		}
		return t.mapper.MapSubClassOnLoad(user, application.Db), nil
	} else {
		err := application.Db.Preload("Login").Find(&user, id).Error
		if err != nil {
			return user, err
		}
		return user, nil
	}
}

func (t *UserRepository) Create(user structs.User) (structs.User, error) {
	err := application.Db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return t.mapper.MapSubClassOnSave(user, application.Db), nil
}
