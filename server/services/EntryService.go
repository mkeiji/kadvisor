package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
	"net/http"
)

type EntryService struct {
	repository repository.EntryRepository
}

func (svc *EntryService) GetManyByUserId(
	userID int,
	limit int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := svc.repository.FindAllByUserId(userID, limit)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (svc *EntryService) GetManyByClassId(
	classID int,
	limit int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := svc.repository.FindAllByClassId(classID, limit)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (svc *EntryService) GetOneById(
	id int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := svc.repository.FindOne(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (svc *EntryService) Post(
	entry structs.Entry,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := svc.repository.Create(entry)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}

func (svc *EntryService) Put(
	entry structs.Entry,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := svc.repository.Update(entry)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}

func (svc *EntryService) Delete(
	id int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := svc.repository.Delete(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}
