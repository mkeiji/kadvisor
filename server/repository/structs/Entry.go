package structs

import "github.com/jinzhu/gorm"

type Entry struct {
	Base
	UserID 			int 	`json:"userID"`
	ClassID 		int		`json:"classID"`
	SubClassID		int		`json:"subClassID"`
	Amount			float64	`json:"amount"`
	Description		string	`json:"description"`
	Obs				string	`json:"obs"`
}

func (e Entry) IsInitializable() bool {return false}

func (e Entry) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Entry{})
}

func (e Entry) Initialize(db *gorm.DB) {}