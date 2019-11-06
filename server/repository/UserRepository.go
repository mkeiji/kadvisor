package repository

import (
	structs2 "kadvisor/server/repository/structs"
	application2 "kadvisor/server/resources/application"
)

type UserRepository struct {}

func (t *UserRepository) FindAll() []structs2.User {
	var products []structs2.User
	application2.Db.Find(&products)
	return products
}

func (t *UserRepository) Create(user structs2.User) structs2.User {
	application2.Db.Create(&user)
	return user
}
