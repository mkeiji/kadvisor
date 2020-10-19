package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type ClassRepository struct{}

func (repo *ClassRepository) FindAllByUserId(
	userID int) ([]structs.Class, error) {

	queryStruct := structs.Class{UserID: userID}
	var classes []structs.Class
	err := application.Db.Where(&queryStruct).Find(&classes).Error

	return classes, err
}

func (repo *ClassRepository) FindOne(
	classID int,
) (structs.Class, error) {
	var class structs.Class
	err := application.Db.Where("id=?", classID).First(&class).Error
	return class, err
}

func (repo *ClassRepository) Create(
	class structs.Class) (structs.Class, error) {
	err := application.Db.Save(&class).Error
	return class, err
}

func (repo *ClassRepository) Update(
	class structs.Class) (structs.Class, error) {
	var stored structs.Class
	err := application.Db.Find(&stored, class.ID).Updates(class).Error
	return stored, err
}

func (repo *ClassRepository) Delete(
	classID int) (int, error) {
	var classToDelete structs.Class
	var err error

	err = application.Db.First(&classToDelete, classID).Error
	if err == nil {
		err = application.Db.Delete(&classToDelete).Error
	}

	return int(classToDelete.ID), err
}
