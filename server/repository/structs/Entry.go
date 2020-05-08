package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Entry struct {
	Base
	UserID 			int 		`json:"userID,omitempty"`
	ClassID 		int			`json:"classID,omitempty"`
	EntryTypeCodeID	string		`json:"entryTypeCodeID,omitempty"` // is lookup
	Date 			time.Time 	`json:"date,omitempty"`
	Amount			float64		`json:"amount,omitempty"`
	Description		string		`json:"description,omitempty"`
	Obs				string		`json:"obs,omitempty"`
}

func (e Entry) IsInitializable() bool {return false}

func (e Entry) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Entry{})
}

func (e Entry) Initialize(db *gorm.DB) {}

/* GORM HOOKS */
func (e *Entry) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&e).Update("is_active", false)
	return
}