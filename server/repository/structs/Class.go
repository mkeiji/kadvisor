package structs

import (
	"gorm.io/gorm"
)

type Class struct {
	Base
	UserID      int    `json:"userID,omitempty"`
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
}

func (e Class) IsInitializable() bool { return false }

func (e Class) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Class{})
}

func (e Class) Initialize(db *gorm.DB) {}
