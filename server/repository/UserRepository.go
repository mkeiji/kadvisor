package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type UserRepository struct{}

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

	return users, nil
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
		return user, nil
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
	return user, nil
}

func (t *UserRepository) Update(user structs.User) (structs.User, error) {
	var stored structs.User
	var err error

	err = application.Db.Set(
		"gorm:association_autocreate", false).Find(
		&stored, user.ID).Updates(user).Error
	return stored, err
}

func (t *UserRepository) Delete(userID int) (structs.User, error) {
	var user structs.User
	var err error

	err = application.Db.First(&user, userID).Error
	if err == nil {
		err = application.Db.Delete(&user).Error
	}

	return user, err
}
