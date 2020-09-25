package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (svc *UserService) GetMany(preloaded bool) ([]structs.User, error) {
	return svc.userRepository.FindAll(preloaded)
}

func (svc *UserService) GetOne(id int, preloaded bool) (structs.User, error) {
	return svc.userRepository.FindOne(id, preloaded)
}

func (svc *UserService) Post(user structs.User) (structs.User, error) {
	return svc.userRepository.Create(user)
}

func (svc *UserService) Put(user structs.User) (structs.User, error) {
	return svc.userRepository.Update(user)
}

func (svc *UserService) Delete(userID int) (structs.User, error) {
	return svc.userRepository.Delete(userID)
}
