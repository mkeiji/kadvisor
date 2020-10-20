package services

import (
	"errors"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
	"net/http"
)

type ClassService struct {
	repository repository.ClassRepository
}

func (svc *ClassService) GetClass(userID int, classID int) dtos.KhttpResponse {
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

func (svc *ClassService) GetManyByUserId(userID int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	classes, err := svc.repository.FindAllByUserId(userID)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, classes)
	}

	return response
}

func (svc *ClassService) GetOneById(id int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	class, err := svc.repository.FindOne(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, class)
	}

	return response
}

func (svc *ClassService) Post(
	class structs.Class,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	class, err := svc.repository.Create(class)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, class)
	}

	return response
}

func (svc *ClassService) Put(
	class structs.Class,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	class, err := svc.repository.Update(class)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, class)
	}

	return response
}

func (svc *ClassService) Delete(
	id int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	class, err := svc.repository.Delete(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, class)
	}

	return response
}
