package repository

import (
	"kadvisor/server/repository/mappers"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type ForecastRepository struct {
	entryMapper mappers.ForecastEntryMapper
}

func (repo ForecastRepository) FindOne(
	userID int, year int, isPreloaded bool,
) (structs.Forecast, error) {
	var forecast structs.Forecast
	var err error

	query := "user_id=? AND year=?"
	if isPreloaded {
		err = application.Db.Preload(
			"Entries").Where(query, userID, year).First(&forecast).Error
	} else {
		err = application.Db.Where(query, userID, year).First(&forecast).Error
	}

	return forecast, err
}

func (repo ForecastRepository) Create(
	forecast structs.Forecast,
) (structs.Forecast, error) {
	var mappedEntries []structs.ForecastEntry
	for _, e := range forecast.Entries {
		mappedEntries = append(mappedEntries, repo.entryMapper.MapForecastEntry(e))
	}
	forecast.Entries = mappedEntries

	err := application.Db.Save(&forecast).Error
	return forecast, err
}

func (repo ForecastRepository) Delete(id int) (structs.Forecast, error) {
	var forecast structs.Forecast
	var err error

	err = application.Db.First(&forecast, id).Error
	if err == nil {
		err = application.Db.Delete(&forecast).Error
	}

	return forecast, err
}
