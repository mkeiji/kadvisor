package validators

import (
	"errors"
	"github.com/jinzhu/gorm"
	"kadvisor/server/repository/structs"
)

type EntryValidator struct {}

func (e *EntryValidator) Validate(
	db *gorm.DB, entry structs.Entry) (err error) {

	var class 		structs.Class
	var subclass 	structs.SubClass

	scErr := db.Find(&subclass, entry.SubClassID).Error
	if scErr != nil {
		err = errors.New("subclass not found")
	}

	cErr := db.Find(&class, entry.ClassID).Error
	if cErr != nil {
		err = errors.New("class not found")
	}

	return
}