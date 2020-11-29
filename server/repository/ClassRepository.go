package repository

import (
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"
)

type ClassRepository struct{}

func (repo ClassRepository) FindAllByUserId(
	userID int,
) ([]s.Class, error) {

	queryStruct := s.Class{UserID: userID}
	var classes []s.Class
	err := app.Db.Where(&queryStruct).Find(&classes).Error

	return classes, err
}

func (repo ClassRepository) FindOne(
	classID int,
) (s.Class, error) {
	class := s.Class{
		Base: s.Base{
			ID: classID,
		},
	}
	err := app.Db.First(&class).Error

	return class, err
}

func (repo ClassRepository) Create(
	class s.Class,
) (s.Class, error) {
	err := app.Db.Save(&class).Error
	return class, err
}

func (repo ClassRepository) Update(
	class s.Class,
) (s.Class, error) {
	stored, err := repo.FindOne(class.ID)
	if err == nil {
		err = app.Db.Model(&stored).Updates(class).Error
	}
	return stored, err
}

func (repo ClassRepository) Delete(
	classID int,
) (s.Class, error) {
	var classToDelete s.Class
	var err error

	err = app.Db.First(&classToDelete, classID).Error
	if err == nil {
		err = app.Db.Delete(&classToDelete).Error
	}

	return classToDelete, err
}
