package interfaces

import s "kadvisor/server/repository/structs"

//go:generate mockgen -destination=mocks/mock_forecast_entry_repository.go -package=mocks . ForecastEntryRepository
type ForecastEntryRepository interface {
	FindOne(id int) (s.ForecastEntry, error)
	Update(entry s.ForecastEntry) (s.ForecastEntry, error)
}
