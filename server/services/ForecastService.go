package services

import (
	"errors"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
	"net/http"
)

type ForecastService struct {
	repository repository.ForecastRepository
}

func (svc *ForecastService) GetOne(
	userID int,
	year int,
	isPreloaded bool,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	if year == 0 {
		yErr := errors.New("year param is required")
		return dtos.NewKresponse(http.StatusBadRequest, yErr)
	}

	forecast, err := svc.repository.FindOne(userID, year, isPreloaded)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, forecast)
	}

	return response
}

func (svc *ForecastService) Post(
	forecast structs.Forecast,
) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	frcast, err := svc.repository.Create(forecast)
	if err != nil {
		response = dtos.NewKresponse(http.StatusBadRequest, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, frcast)
	}

	return response
}

func (svc *ForecastService) Delete(id int) dtos.KhttpResponse {
	var response dtos.KhttpResponse

	frcast, err := svc.repository.Delete(id)
	if err != nil {
		response = dtos.NewKresponse(http.StatusNotFound, err)
	} else {
		response = dtos.NewKresponse(http.StatusOK, frcast)
	}

	return response
}
