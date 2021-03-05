package repository

import (
	"kadvisor/server/repository/mappers"
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"

	"gorm.io/gorm"
)

type EntryRepository struct {
	Db     *gorm.DB
	Mapper mappers.EntryMapper
}

func NewEntryRepository() EntryRepository {
	return EntryRepository{
		Db:     app.Db,
		Mapper: mappers.EntryMapper{},
	}
}

func (this EntryRepository) FindAllByUserId(
	userID int, limit int,
) ([]s.Entry, error) {
	queryStruct := s.Entry{UserID: userID}
	return this.getEntries(queryStruct, limit)
}

func (this EntryRepository) FindAllByClassId(
	classID int, limit int) ([]s.Entry, error) {

	queryStruct := s.Entry{ClassID: classID}
	return this.getEntries(queryStruct, limit)
}

func (this EntryRepository) FindOne(id int) (s.Entry, error) {
	var entry s.Entry
	err := this.Db.Where("id=?", id).First(&entry).Error
	return this.Mapper.MapEntryDate(entry), err
}

func (this EntryRepository) Create(
	entry s.Entry,
) (s.Entry, error) {
	eMapped := this.Mapper.MapEntry(entry)
	err := this.Db.Save(&eMapped).Error
	return eMapped, err
}

func (this EntryRepository) Update(
	entry s.Entry,
) (s.Entry, error) {
	eMapped := this.Mapper.MapEntry(entry)
	stored, err := this.FindOne(entry.ID)
	if err == nil {
		err = this.Db.Model(&stored).Updates(eMapped).Error
		if entry.Amount == 0 {
			err = this.Db.Model(&stored).
				UpdateColumn("amount", 0).
				Error
		}
	}
	return stored, err
}

func (this EntryRepository) Delete(id int) (int, error) {
	var entry s.Entry
	var err error

	err = this.Db.First(&entry, id).Error
	if err == nil {
		err = this.Db.Delete(&entry).Error
	}

	return entry.ID, err
}

func (this EntryRepository) getEntries(query s.Entry, limit int) ([]s.Entry, error) {
	var entries []s.Entry
	var err error

	dbQuery := this.Db.Order("created_at desc")
	if limit > 0 {
		err = dbQuery.Limit(limit).Where(query).Find(&entries).Error
	} else {
		err = dbQuery.Where(query).Find(&entries).Error
	}

	return this.Mapper.MapEntriesDates(entries), err
}
