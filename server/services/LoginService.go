package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
	"net/http"
)

type LoginService struct {
	loginRepository repository.LoginRepository
}

func (l *LoginService) GetOneByEmail(email string) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	login, err := l.loginRepository.FindOneByEmail(email)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, login)
	}

	return response
}

func (l *LoginService) Put(login structs.Login) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	login, err := l.loginRepository.Update(login)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, login)
	}

	return response
}

func (l *LoginService) UpdateLoginStatus(
	login structs.Login,
	isLoggedIn bool,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	login, err := l.loginRepository.UpdateLoginStatus(login, isLoggedIn)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, login)
	}

	return response
}
