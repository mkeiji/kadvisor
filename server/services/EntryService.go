package services

import (
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
	"net/http"

	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
)

type EntryService struct {
	Repository i.EntryRepository
}

func NewEntryService() EntryService {
	return EntryService{
		Repository: r.EntryRepository{},
	}
}

func (svc EntryService) GetManyByUserId(
	userID int,
	limit int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := svc.Repository.FindAllByUserId(userID, limit)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (svc EntryService) GetManyByClassId(
	classID int,
	limit int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := svc.Repository.FindAllByClassId(classID, limit)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (svc EntryService) GetOneById(
	id int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := svc.Repository.FindOne(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (svc EntryService) Post(
	entry s.Entry,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := svc.Repository.Create(entry)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}

func (svc EntryService) Put(
	entry s.Entry,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := svc.Repository.Update(entry)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}

func (svc EntryService) Delete(
	id int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := svc.Repository.Delete(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}
