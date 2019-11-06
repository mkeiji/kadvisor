package interfaces

import "github.com/jinzhu/gorm"

type Entity interface {
	InitializeTable(db *gorm.DB)
}

