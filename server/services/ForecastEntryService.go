package services

import (
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
	"net/http"
)

type ForecastEntryService struct {
	repository repository.ForecastEntryRepository
}

func (svc *ForecastEntryService) Put(
	entry structs.ForecastEntry,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	updated, err := svc.repository.Update(entry)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, updated)
	}

	return response
}
