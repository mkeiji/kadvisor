package repository

import (
	"kadvisor/server/repository/mappers"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type ForecastRepository struct {
	entryMapper 	mappers.ForecastEntryMapper
}

func (repo *ForecastRepository) FindForecastByUser(
	userID int, isPreloaded bool) (structs.Forecast, error) {

	var forecast structs.Forecast
	var err error

	if isPreloaded {
		err = application.Db.Preload(
			"Entries").Where("user_id=?", userID).Find(&forecast).Error
	} else {
		err = application.Db.Where("user_id=?", userID).Find(&forecast).Error
	}

	return forecast, err
}

func (repo *ForecastRepository) Create(
	forecast structs.Forecast) (structs.Forecast, error) {

	var mappedEntries []structs.ForecastEntry
	for _, e := range forecast.Entries {
		mappedEntries = append(mappedEntries, repo.entryMapper.MapForecastEntry(e))
	}
	forecast.Entries = mappedEntries
	
	err := application.Db.Save(&forecast).Error
	return forecast, err
}

func (repo *ForecastRepository) Delete(id int) (structs.Forecast, error) {
	forecast := structs.Forecast{Base: structs.Base{ID: id}}
	err := application.Db.Delete(&forecast).Error
	return forecast, err
}