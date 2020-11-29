package interfaces

import (
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
)

//go:generate mockgen -destination=mocks/mock_forecast_entry_service.go -package=mocks . ForecastEntryService
type ForecastEntryService interface {
	Put(entry s.ForecastEntry) dtos.KhttpResponse
}
