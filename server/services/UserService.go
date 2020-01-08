package services

import (
	"kadvisor/server/repository"
	structs2 "kadvisor/server/repository/structs"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (t *UserService) GetMany() []structs2.User {
	return t.userRepository.FindAll()
}

func (t *UserService) GetOne(id int) structs2.User {
	return t.userRepository.FindOne(id)
}

func (t *UserService) Post(user structs2.User) structs2.User {
	return t.userRepository.Create(user)
}