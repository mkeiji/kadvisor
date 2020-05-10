package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type ForecastEntryRepository struct {}

func (repo *ForecastEntryRepository) Update(
	entry structs.ForecastEntry) (structs.ForecastEntry, error) {

	var stored structs.ForecastEntry
	err := application.Db.Find(&stored, entry.ID).Updates(entry).Error
	return  stored, err
}