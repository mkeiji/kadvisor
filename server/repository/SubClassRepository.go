package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type SubClassRepository struct {}

func (rep *SubClassRepository) FindAllByUserId(
	userID int) ([]structs.SubClass, error) {

	queryStruct := structs.SubClass{UserID: userID}
	return findManyByStructQuery(queryStruct)
}

func (rep *SubClassRepository) FindAllByClassId(
	classID int) ([]structs.SubClass, error) {

	queryStruct := structs.SubClass{ClassID: classID}
	return findManyByStructQuery(queryStruct)
}

func (rep *SubClassRepository) FindOne(
	id int) (structs.SubClass, error) {

	var subclass structs.SubClass

	err := application.Db.Find(
		&subclass, id).Error
	return subclass, err
}

func (rep *SubClassRepository) Create(
	subclass structs.SubClass) (structs.SubClass, error) {
	err := application.Db.Save(&subclass).Error
	return subclass, err
}

func (rep *SubClassRepository) Delete(
	id int) (int, error) {
	toDelete := structs.SubClass{Base: structs.Base{ID: uint(id)}}
	err := application.Db.Delete(&toDelete).Error
	return int(toDelete.ID), err
}

func findManyByStructQuery(
	queryStruct structs.SubClass) ([]structs.SubClass, error) {
	var subclasses []structs.SubClass

	err := application.Db.Where(
		&queryStruct).Find(
		&subclasses).Error
	return subclasses, err
}