package interfaces

import (
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
)

//go:generate mockgen -destination=mocks/mock_forecast_service.go -package=mocks . ForecastService
type ForecastService interface {
	GetOne(userID int, year int, isPreloaded bool) dtos.KhttpResponse
	Post(forecast s.Forecast) dtos.KhttpResponse
	Delete(id int) dtos.KhttpResponse
}
