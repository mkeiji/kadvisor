package services

import (
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
	"net/http"

	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
)

type LoginService struct {
	Repository i.LoginRepository
}

func NewLoginService() LoginService {
	return LoginService{
		Repository: r.NewLoginRepository(),
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

func (l LoginService) Put(login s.Login) dtos.KhttpResponse {
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
	login s.Login,
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
