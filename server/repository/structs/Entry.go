package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Entry struct {
	Base
	UserID 			int 		`json:"userID"`
	ClassID 		int			`json:"classID"`
	SubClassID		int			`json:"subClassID"`
	Date 			time.Time 	`json:"date"`
	Amount			float64		`json:"amount"`
	Description		string		`json:"description"`
	Obs				string		`json:"obs"`
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