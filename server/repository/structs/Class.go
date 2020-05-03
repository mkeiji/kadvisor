package structs

import (
	"github.com/jinzhu/gorm"
	"os"
)

type Class struct {
	Base
	UserID 		int 		`json:"userID,omitempty"`
	Name 		string 		`json:"name,omitempty"`
	Description	string 		`json:"description,omitempty"`
	SubClasses	[]SubClass	`gorm:"ForeignKey:ClassID" json:"subClasses,omitempty"`
}
func (e Class) IsInitializable() bool {return false}

func (e Class) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == os.Getenv("DEV_ENV") {
		db.DropTableIfExists(&SubClass{})
	}
	db.AutoMigrate(&Class{})
}

func (e Class) Initialize(db *gorm.DB) {}

/* GORM HOOKS */
func (e *Class) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&e).Update("is_active", false)
	return
}