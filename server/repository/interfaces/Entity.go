package interfaces

import "github.com/jinzhu/gorm"

type Entity interface {
	Migrate(db *gorm.DB)
	IsInitializable() bool
	Initialize(db *gorm.DB)
}

