package services

import (
	"kadvisor/server/libs/dtos"
	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
	s "kadvisor/server/repository/structs"
	"net/http"
)

type UserService struct {
	Repository i.UserRepository
}

func NewUserService() UserService {
	return UserService{
		Repository: r.UserRepository{},
	}
}

func (svc UserService) GetMany(preloaded bool) dtos.KhttpResponse {
	var res dtos.KhttpResponse
	users, err := svc.Repository.FindAll(preloaded)
	if err != nil {
		res = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, users)
	}
	return res
}

func (svc UserService) GetOne(id int, preloaded bool) dtos.KhttpResponse {
	var res dtos.KhttpResponse

	user, err := svc.Repository.FindOne(id, preloaded)
	if err != nil {
		res = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, user)
	}

	return res
}

func (svc UserService) Post(user s.User) dtos.KhttpResponse {
	var res dtos.KhttpResponse

	user, err := svc.Repository.Create(user)
	if err != nil {
		res = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, user)
	}

	return res
}

func (svc UserService) Put(user s.User) dtos.KhttpResponse {
	var res dtos.KhttpResponse

	user, err := svc.Repository.Update(user)
	if err != nil {
		res = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, user)
	}

	return res
}

func (svc UserService) Delete(userID int) dtos.KhttpResponse {
	var res dtos.KhttpResponse

	user, err := svc.Repository.Delete(userID)
	if err != nil {
		res = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, user)
	}

	return res
}
