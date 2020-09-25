package interfaces

import "gorm.io/gorm"

type Validator interface {
	Validate(entity Entity, db *gorm.DB)
}
