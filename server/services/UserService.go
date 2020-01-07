package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (t *UserService) GetMany() []structs.User {
	return t.userRepository.FindAll()
}

func (t *UserService) GetOne(id int) structs.User {
	return t.userRepository.FindOne(id)
}

func (t *UserService) Post(user structs.User) structs.User {
	return t.userRepository.Create(user)
}