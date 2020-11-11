package repository

import (
	"kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
)

type EntryRepository struct {
	mapper mappers.EntryMapper
}

func (repo EntryRepository) FindAllByUserId(
	userID int, limit int) ([]s.Entry, error) {

	queryStruct := s.Entry{UserID: userID}
	return getEntries(queryStruct, limit)
}

func (repo EntryRepository) FindAllByClassId(
	classID int, limit int) ([]s.Entry, error) {

	queryStruct := s.Entry{ClassID: classID}
	return getEntries(queryStruct, limit)
}

func (repo EntryRepository) FindOne(id int) (s.Entry, error) {
	var entry s.Entry
	err := app.Db.Where("id=?", id).First(&entry).Error
	return entry, err
}

func (repo EntryRepository) Create(
	entry s.Entry,
) (s.Entry, error) {
	eMapped := repo.mapper.MapEntry(entry)
	err := app.Db.Save(&eMapped).Error
	return eMapped, err
}

func (repo EntryRepository) Update(
	entry s.Entry,
) (s.Entry, error) {
	eMapped := repo.mapper.MapEntry(entry)
	stored, err := repo.FindOne(entry.ID)
	if err == nil {
		err = app.Db.Model(&stored).Updates(eMapped).Error
		if entry.Amount == 0 {
			err = app.Db.Model(&stored).
				UpdateColumn("amount", 0).
				Error
		}
	}
	return stored, err
}

func (repo EntryRepository) Delete(id int) (int, error) {
	var entry s.Entry
	var err error

	err = app.Db.First(&entry, id).Error
	if err == nil {
		err = app.Db.Delete(&entry).Error
	}

	return entry.ID, err
}

func getEntries(query s.Entry, limit int) ([]s.Entry, error) {
	var entries []s.Entry
	var err error

	dbQuery := app.Db.Order("created_at desc")
	if limit > 0 {
		err = dbQuery.Limit(limit).Where(query).Find(&entries).Error
	} else {
		err = dbQuery.Where(query).Find(&entries).Error
	}

	return entries, err
}
