package structs

import (
	"gorm.io/gorm"
	"time"
)

type Entry struct {
	Base
	UserID          int       `json:"userID,omitempty" validate:"required"`
	ClassID         int       `json:"classID,omitempty" validate:"required,ispositive"`
	EntryTypeCodeID string    `json:"entryTypeCodeID,omitempty"` // is lookup
	Date            time.Time `json:"date,omitempty"`
	Amount          float64   `json:"amount,omitempty"`
	Description     string    `json:"description,omitempty"`
	Obs             string    `json:"obs,omitempty"`
}

func (e Entry) IsInitializable() bool { return false }

func (e Entry) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Entry{})
}

func (e Entry) Initialize(db *gorm.DB) {}
