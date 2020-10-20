package repository

import (
	"errors"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
)

type UserRepository struct{}

func (t *UserRepository) FindAll(preloaded bool) ([]s.User, error) {
	var users []s.User
	var err error

	if preloaded {
		err = app.Db.
			Preload("Login").
			Preload("Entries").
			Preload("Classes").
			Find(&users).
			Error
	} else {
		err = app.Db.Preload("Login").Find(&users).Error
	}

	return t.handleFindManyErr(users, err)
}

func (t *UserRepository) FindOne(id int, preloaded bool) (s.User, error) {
	var user s.User
	var err error

	if preloaded {
		err = app.Db.
			Preload("Login").
			Preload("Entries").
			Preload("Classes").
			First(&user, id).
			Error
	} else {
		err = app.Db.Preload("Login").First(&user, id).Error
	}

	return user, err
}

func (t *UserRepository) Create(user s.User) (s.User, error) {
	err := app.Db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (t *UserRepository) Update(user s.User) (s.User, error) {
	var stored s.User
	var err error

	err = app.Db.Set(
		"gorm:association_autocreate", false).First(
		&stored, user.ID).Updates(user).Error
	return stored, err
}

func (t *UserRepository) Delete(userID int) (s.User, error) {
	var user s.User
	var err error

	err = app.Db.First(&user, userID).Error
	if err == nil {
		err = app.Db.Delete(&user).Error
	}

	return user, err
}

func (t *UserRepository) handleFindManyErr(
	users []s.User,
	err error,
) ([]s.User, error) {
	if len(users) < 1 {
		err = errors.New("records not found")
	}
	return users, err
}
