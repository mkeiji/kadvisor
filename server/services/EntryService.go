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
		Repository: r.NewEntryRepository(),
	}
}

func (this EntryService) GetManyByUserId(
	userID int,
	limit int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := this.Repository.FindAllByUserId(userID, limit)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (this EntryService) GetManyByClassId(
	classID int,
	limit int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := this.Repository.FindAllByClassId(classID, limit)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (this EntryService) GetOneById(
	id int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	entries, err := this.Repository.FindOne(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, entries)
	}

	return response
}

func (this EntryService) Post(
	entry s.Entry,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := this.Repository.Create(entry)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}

func (this EntryService) Put(
	entry s.Entry,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := this.Repository.Update(entry)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}

func (this EntryService) Delete(
	id int,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	nEntry, err := this.Repository.Delete(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, nEntry)
	}

	return response
}
