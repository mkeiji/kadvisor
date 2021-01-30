package repository

import (
	"kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"

	"gorm.io/gorm"
)

type ForecastEntryRepository struct {
	Mapper mappers.ForecastEntryMapper
	Db     *gorm.DB
}

func NewForecastEntryRepository() ForecastEntryRepository {
	return ForecastEntryRepository{
		Mapper: mappers.ForecastEntryMapper{},
		Db:     app.Db,
	}
}

func (this ForecastEntryRepository) FindOne(id int) (s.ForecastEntry, error) {
	var entry s.ForecastEntry
	err := this.Db.Where("id=?", id).First(&entry).Error
	return entry, err
}

func (this ForecastEntryRepository) Update(
	entry s.ForecastEntry,
) (s.ForecastEntry, error) {
	eMapped := this.Mapper.MapForecastEntry(entry)

	stored, err := this.FindOne(entry.ID)
	if err == nil {
		err = this.Db.Model(&stored).
			Select("income", "expense").
			UpdateColumns(
				s.ForecastEntry{
					Income:  eMapped.Income,
					Expense: eMapped.Expense,
				},
			).Error
	}

	return stored, err
}
