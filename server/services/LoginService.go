package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	"net/http"

	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
)

type LoginService struct {
	Repository i.LoginRepository
}

func NewLoginService() LoginService {
	return LoginService{
		Repository: r.LoginRepository{},
	}
}

func (l LoginService) GetOneByEmail(email string) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	login, err := l.Repository.FindOneByEmail(email)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, login)
	}

	return response
}

func (l LoginService) Put(login structs.Login) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	login, err := l.Repository.Update(login)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, login)
	}

	return response
}

func (l LoginService) UpdateLoginStatus(
	login structs.Login,
	isLoggedIn bool,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	login, err := l.Repository.UpdateLoginStatus(login, isLoggedIn)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, login)
	}

	return response
}
