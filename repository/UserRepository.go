package repository

import (
	"kadvisor/repository/structs"
	"kadvisor/resources/application"
)

type UserRepository struct {}

func (t *UserRepository) FindAll() []structs.User {
	var products []structs.User
	application.Db.Find(&products)
	return products
}

func (t *UserRepository) Create(user structs.User) structs.User {
	application.Db.Create(&user)
	return user
}
