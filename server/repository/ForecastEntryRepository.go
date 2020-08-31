package repository

import (
	"kadvisor/server/repository/mappers"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type ForecastEntryRepository struct {
	mapper mappers.ForecastEntryMapper
}

func (repo *ForecastEntryRepository) Update(
	entry structs.ForecastEntry) (structs.ForecastEntry, error) {

	var stored structs.ForecastEntry
	eMapped := repo.mapper.MapForecastEntry(entry)

	err := application.Db.Find(&stored, entry.ID).Updates(eMapped).Error
	return stored, err
}
