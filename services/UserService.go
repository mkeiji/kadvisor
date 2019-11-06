package services

import (
	"kadvisor/repository"
	"kadvisor/repository/structs"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (t *UserService) GetMany() []structs.User {
	return t.userRepository.FindAll()
}

func (t *UserService) Post(user structs.User) structs.User {
	return t.userRepository.Create(user)
}