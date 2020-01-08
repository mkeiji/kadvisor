package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type UserRepository struct {}

func (t *UserRepository) FindAll() []structs.User {
	var users []structs.User
	application.Db.Preload("Login").Find(&users)
	return users
}

func (t *UserRepository) FindOne(id int) structs.User {
	var user structs.User
	application.Db.Preload("Login").Find(&user, id)
	return user
}

func (t *UserRepository) Create(user structs.User) structs.User {
	application.Db.Save(&user)
	return user
}
