package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
	"net/http"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (svc *UserService) GetMany(preloaded bool) dtos.KhttpResponse {
	var res dtos.KhttpResponse
	users, err := svc.userRepository.FindAll(preloaded)
	if err != nil {
		res = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, users)
	}
	return res
}

func (svc *UserService) GetOne(id int, preloaded bool) dtos.KhttpResponse {
	var res dtos.KhttpResponse

	user, err := svc.userRepository.FindOne(id, preloaded)
	if err != nil {
		res = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, user)
	}

	return res
}

func (svc *UserService) Post(user structs.User) dtos.KhttpResponse {
	var res dtos.KhttpResponse

	user, err := svc.userRepository.Create(user)
	if err != nil {
		res = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, user)
	}

	return res
}

func (svc *UserService) Put(user structs.User) dtos.KhttpResponse {
	var res dtos.KhttpResponse

	user, err := svc.userRepository.Update(user)
	if err != nil {
		res = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, user)
	}

	return res
}

func (svc *UserService) Delete(userID int) dtos.KhttpResponse {
	var res dtos.KhttpResponse

	user, err := svc.userRepository.Delete(userID)
	if err != nil {
		res = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		res = dtos.NewKresponse(http.StatusOK, user)
	}

	return res
}
