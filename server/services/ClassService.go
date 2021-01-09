package services

import (
	"errors"
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
	"net/http"

	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
)

type ClassService struct {
	Repository i.ClassRepository
}

func NewClassService() ClassService {
	return ClassService{
		Repository: r.NewClassRepository(),
	}
}

func (svc ClassService) GetClass(userID int, classID int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	shouldGetClassesByUserId := classID == 0 && userID != 0
	if classID != 0 {
		response = svc.GetOneById(classID)
	} else if shouldGetClassesByUserId {
		response = svc.GetManyByUserId(userID)
	} else {
		response = dtos.NewKresponse(
			http.StatusBadRequest,
			errors.New("query param error"),
		)
	}

	return response
}

func (svc ClassService) GetManyByUserId(userID int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	classes, err := svc.Repository.FindAllByUserId(userID)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, classes)
	}

	return response
}

func (svc ClassService) GetOneById(id int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	class, err := svc.Repository.FindOne(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, class)
	}

	return response
}

func (svc ClassService) Post(
	class s.Class,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	class, err := svc.Repository.Create(class)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, class)
	}

	return response
}

func (svc ClassService) Put(
	class s.Class,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	class, err := svc.Repository.Update(class)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, class)
	}

	return response
}

func (svc ClassService) Delete(
	id int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	class, err := svc.Repository.Delete(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, class)
	}

	return response
}
