package repository

import (
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"

	"gorm.io/gorm"
)

type ClassRepository struct {
	Db *gorm.DB
}

func NewClassRepository() ClassRepository {
	return ClassRepository{Db: app.Db}
}

func (this ClassRepository) FindAllByUserId(
	userID int,
) ([]s.Class, error) {

	queryStruct := s.Class{UserID: userID}
	var classes []s.Class
	err := this.Db.Where(&queryStruct).Find(&classes).Error

	return classes, err
}

func (this ClassRepository) FindOne(
	classID int,
) (s.Class, error) {
	class := s.Class{
		Base: s.Base{
			ID: classID,
		},
	}
	err := this.Db.First(&class).Error

	return class, err
}

func (this ClassRepository) Create(
	class s.Class,
) (s.Class, error) {
	err := this.Db.Save(&class).Error
	return class, err
}

func (this ClassRepository) Update(
	class s.Class,
) (s.Class, error) {
	stored, err := this.FindOne(class.ID)
	if err == nil {
		err = this.Db.Model(&stored).Updates(class).Error
	}
	return stored, err
}

func (this ClassRepository) Delete(
	classID int,
) (s.Class, error) {
	var classToDelete s.Class
	var err error

	err = this.Db.First(&classToDelete, classID).Error
	if err == nil {
		err = this.Db.Delete(&classToDelete).Error
	}

	return classToDelete, err
}
