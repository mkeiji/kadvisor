package repository

import (
	"kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"

	"gorm.io/gorm"
)

type ForecastRepository struct {
	EntryMapper mappers.ForecastEntryMapper
	Db          *gorm.DB
}

func NewForecastRepository() ForecastRepository {
	return ForecastRepository{
		EntryMapper: mappers.ForecastEntryMapper{},
		Db:          app.Db,
	}
}

func (this ForecastRepository) FindOne(
	userID int, year int, isPreloaded bool,
) (s.Forecast, error) {
	var forecast s.Forecast
	var err error

	query := "user_id=? AND year=?"
	if isPreloaded {
		err = this.Db.Preload("Entries").
			Where(query, userID, year).
			First(&forecast).Error
	} else {
		err = this.Db.
			Where(query, userID, year).
			First(&forecast).Error
	}

	return forecast, err
}

func (this ForecastRepository) Create(
	forecast s.Forecast,
) (s.Forecast, error) {
	var mappedEntries []s.ForecastEntry
	for _, e := range forecast.Entries {
		mappedEntries = append(
			mappedEntries,
			this.EntryMapper.MapForecastEntry(e),
		)
	}
	forecast.Entries = mappedEntries

	err := this.Db.Save(&forecast).Error
	return forecast, err
}

func (this ForecastRepository) Delete(id int) (s.Forecast, error) {
	var forecast s.Forecast
	var err error

	err = this.Db.First(&forecast, id).Error
	if err == nil {
		err = this.Db.Delete(&forecast).Error
	}

	return forecast, err
}
