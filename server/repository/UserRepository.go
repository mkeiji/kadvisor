package repository

import (
	"errors"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"

	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository() UserRepository {
	return UserRepository{
		Db: app.Db,
	}
}

func (this UserRepository) FindAll(preloaded bool) ([]s.User, error) {
	var users []s.User
	var err error

	if preloaded {
		err = this.Db.
			Preload("Login").
			Preload("Entries").
			Preload("Classes").
			Find(&users).
			Error
	} else {
		err = this.Db.Preload("Login").Find(&users).Error
	}

	return this.handleFindManyErr(users, err)
}

func (this UserRepository) FindOne(id int, preloaded bool) (s.User, error) {
	var user s.User
	var err error

	if preloaded {
		err = this.Db.
			Preload("Login").
			Preload("Entries").
			Preload("Classes").
			First(&user, id).
			Error
	} else {
		err = this.Db.Preload("Login").First(&user, id).Error
	}

	return user, err
}

func (this UserRepository) Create(user s.User) (s.User, error) {
	err := this.Db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (this UserRepository) Update(user s.User) (s.User, error) {
	stored, err := this.FindOne(user.ID, false)
	if err == nil {
		err = this.Db.Model(&stored).Updates(user).Error
	}

	return stored, err
}

func (this UserRepository) Delete(userID int) (s.User, error) {
	var user s.User
	var err error

	err = this.Db.First(&user, userID).Error
	if err == nil {
		err = this.Db.Delete(&s.User{}, userID).Error
	}

	return user, err
}

func (this UserRepository) handleFindManyErr(
	users []s.User,
	err error,
) ([]s.User, error) {
	if err == nil {
		if len(users) < 1 {
			err = errors.New("records not found")
		}
	}
	return users, err
}
