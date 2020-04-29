package interfaces

import "github.com/jinzhu/gorm"

type Validator interface {
	Validate(entity Entity, db *gorm.DB)
}