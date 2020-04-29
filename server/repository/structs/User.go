package structs

import (
	"github.com/jinzhu/gorm"
	"os"
)

type User struct {
	Base
	FirstName 	string 	`json:"firstName"`
	LastName  	string 	`json:"lastName"`
	Phone     	string 	`json:"phone"`
	Address   	string 	`json:"address"`
	Login     	Login  	`gorm:"ForeignKey:UserID" json:"login"`
	Entries		[]Entry	`gorm:"ForeignKey:UserID" json:"entries"`
	Classes		[]Class	`gorm:"ForeignKey:UserID" json:"classes"`
}
func (e User) IsInitializable() bool { return false }

func (e User) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == os.Getenv("DEV_ENV") {
		db.DropTableIfExists(&Class{})
		db.DropTableIfExists(&Entry{})
		db.DropTableIfExists(&Login{})
		db.DropTableIfExists(&User{})
	}
	db.AutoMigrate(&User{})
}

func (e User) Initialize(db *gorm.DB) {/* empty */}

/* GORM HOOKS */
func (e *User) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&e).Update("is_active", false)
	return
}