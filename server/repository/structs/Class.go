package structs

import (
	"github.com/jinzhu/gorm"
	"os"
)

type Class struct {
	Base
	UserID 		int 		`json:"userID"`
	Name 		string 		`json:"name"`
	Description	string 		`json:"description"`
	SubClasses	[]SubClass	`gorm:"ForeignKey:ClassID" json:"subClasses"`
}
func (e Class) IsInitializable() bool {return false}

func (e Class) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == os.Getenv("DEV_ENV") {
		db.DropTableIfExists(&SubClass{})
	}
	db.AutoMigrate(&Class{})
}

func (e Class) Initialize(db *gorm.DB) {}