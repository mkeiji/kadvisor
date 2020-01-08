package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type LoginService struct {
	loginRepository repository.LoginRepository
}

func (l *LoginService) GetOneByEmail(email string) structs.Login {
	return l.loginRepository.FindOneByEmail(email)
}