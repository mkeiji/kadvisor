package validators

import (
	"errors"
	"github.com/jinzhu/gorm"
	"kadvisor/server/repository/structs"
)

type EntryValidator struct{}

func (e *EntryValidator) Validate(
	db *gorm.DB, entry structs.Entry) (err error) {

	var class structs.Class
	cErr := db.Find(&class, entry.ClassID).Error
	if cErr != nil {
		err = errors.New("class not found")
	}

	var lookup structs.Code
	lErr := db.Where("code_type_id=?", entry.EntryTypeCodeID).Find(&lookup).Error
	if lErr != nil {
		err = errors.New("invalid lookup @ entryTypeCodeID")
	}

	return
}
