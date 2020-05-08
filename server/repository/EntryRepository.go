package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/repository/validators"
	"kadvisor/server/resources/application"
	"time"
)

type EntryRepository struct {
	validator 	validators.EntryValidator
}

func (repo *EntryRepository) FindAllByUserId(
	userID int, limit int) ([]structs.Entry, error) {

	queryStruct := structs.Entry{UserID: userID}
	return getEntries(queryStruct, limit)
}

func (repo *EntryRepository) FindAllByClassId(
	classID int, limit int) ([]structs.Entry, error) {

	queryStruct := structs.Entry{ClassID: classID}
	return getEntries(queryStruct, limit)
}

func (repo *EntryRepository) FindOne(id int) (structs.Entry, error) {
	var entry structs.Entry
	err := application.Db.Find(&entry, id).Error
	return entry, err
}

func (repo *EntryRepository) Create(
	entry structs.Entry) (structs.Entry, error) {

	utc, _ := time.LoadLocation("UTC")
	entry.Date = entry.Date.In(utc)

	vErr := repo.validator.Validate(application.Db, entry)
	if vErr != nil {
		return structs.Entry{}, vErr
	} else {
		err := application.Db.Save(&entry).Error
		return entry, err
	}
}

func (repo *EntryRepository) Update(
	entry structs.Entry) (structs.Entry, error) {
	var stored structs.Entry

	err := application.Db.Find(&stored, entry.ID).Updates(entry).Error
	return stored, err
}

func (repo *EntryRepository) Delete(id int) (int, error) {
	entry := structs.Entry{Base: structs.Base{ID: uint(id)}}
	err := application.Db.Delete(&entry).Error
	return int(entry.ID), err
}

func getEntries(query structs.Entry, limit int) ([]structs.Entry, error){
	var entries []structs.Entry
	var err error

	dbQuery := application.Db.Order("created_at desc")
	if limit > 0 {
		err = dbQuery.Limit(limit).Where(query).Find(&entries).Error
	} else {
		err = dbQuery.Where(query).Find(&entries).Error
	}

	return entries, err
}