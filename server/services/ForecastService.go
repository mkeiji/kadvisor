package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type ForecastService struct {
	repository 	repository.ForecastRepository
}

func (svc *ForecastService) GetOne(
	userID int, year int, isPreloaded bool) (structs.Forecast, error) {
	return svc.repository.FindOne(userID, year, isPreloaded)
}

func (svc *ForecastService) Post (
	forecast structs.Forecast) (structs.Forecast, error) {
	return svc.repository.Create(forecast)
}

func (svc *ForecastService) Delete (id int) (structs.Forecast, error) {
	return svc.repository.Delete(id)
}