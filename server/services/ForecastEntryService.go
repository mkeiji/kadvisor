package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	"net/http"

	r "kadvisor/server/repository"
	i "kadvisor/server/repository/interfaces"
)

type ForecastEntryService struct {
	Repository i.ForecastEntryRepository
}

func NewForecastEntryService() ForecastEntryService {
	return ForecastEntryService{
		Repository: r.ForecastEntryRepository{},
	}
}

func (svc ForecastEntryService) Put(
	entry structs.ForecastEntry,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	updated, err := svc.Repository.Update(entry)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, updated)
	}

	return response
}
