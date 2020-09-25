package interfaces

import "gorm.io/gorm"

type Entity interface {
	Migrate(db *gorm.DB)
	IsInitializable() bool
	Initialize(db *gorm.DB)
}
