package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (t *UserService) GetMany(preloaded bool) ([]structs.User, error) {
	return t.userRepository.FindAll(preloaded)
}

func (t *UserService) GetOne(id int, preloaded bool) (structs.User, error) {
	return t.userRepository.FindOne(id, preloaded)
}

func (t *UserService) Post(user structs.User) (structs.User, error) {
	return t.userRepository.Create(user)
}