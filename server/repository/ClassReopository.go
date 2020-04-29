package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type ClassRepository struct {}

func (repo *ClassRepository) FindAllByUserId(
	userID int, preloaded bool) ([]structs.Class, error) {

	queryStruct := structs.Class{UserID: userID}
	var classes []structs.Class
	var err error

	if preloaded {
		err = application.Db.Where(&queryStruct).Preload(
			"SubClasses").Find(&classes).Error
	} else {
		err = application.Db.Where(&queryStruct).Find(&classes).Error
	}
	return classes, err
}

func (repo *ClassRepository) FindOne(classID int, preloaded bool) (structs.Class, error) {
	var class structs.Class
	var err error

	if preloaded {
		err = application.Db.Preload(
			"SubClasses").Find(&class, classID).Error
	} else {
		err = application.Db.Find(&class, classID).Error
	}
	return class, err
}

func (repo *ClassRepository) Create(
	class structs.Class) (structs.Class, error) {
	err := application.Db.Save(&class).Error
	return class, err
}

func (repo *ClassRepository) Delete(
	classID int) (int, error) {
	classToDelete := structs.Class{
		Base: structs.Base{ID: uint(classID)}}
	err := application.Db.Delete(&classToDelete).Error
	return int(classToDelete.ID), err
}