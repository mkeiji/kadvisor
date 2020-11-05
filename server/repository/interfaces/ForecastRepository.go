package interfaces

import s "kadvisor/server/repository/structs"

//go:generate mockgen -destination=mocks/mock_forecast_repository.go -package=mocks . ForecastRepository
type ForecastRepository interface {
	FindOne(userID int, year int, isPreloaded bool) (s.Forecast, error)
	Create(forecast s.Forecast) (s.Forecast, error)
	Delete(id int) (s.Forecast, error)
}
