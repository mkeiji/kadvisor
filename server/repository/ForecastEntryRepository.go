package repository

import (
	"kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
)

type ForecastEntryRepository struct {
	mapper mappers.ForecastEntryMapper
}

func (repo *ForecastEntryRepository) FindOne(id int) (s.ForecastEntry, error) {
	var entry s.ForecastEntry
	err := app.Db.Where("id=?", id).First(&entry).Error
	return entry, err
}

func (repo *ForecastEntryRepository) Update(
	entry s.ForecastEntry,
) (s.ForecastEntry, error) {
	eMapped := repo.mapper.MapForecastEntry(entry)

	stored, err := repo.FindOne(entry.ID)
	err = app.Db.Model(&stored).Select("income", "expense").
		UpdateColumns(
			s.ForecastEntry{
				Income:  eMapped.Income,
				Expense: eMapped.Expense,
			},
		).Error

	return stored, err
}
