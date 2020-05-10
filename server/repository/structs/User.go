package structs

import (
	"github.com/jinzhu/gorm"
	"os"
)

type User struct {
	Base
	FirstName 	string 		`json:"firstName,omitempty"`
	LastName  	string 		`json:"lastName,omitempty"`
	Phone     	string 		`json:"phone,omitempty"`
	Address   	string 		`json:"address,omitempty"`
	Login     	Login  		`gorm:"ForeignKey:UserID" json:"login,omitempty"`
	Entries		[]Entry		`gorm:"ForeignKey:UserID" json:"entries,omitempty"`
	Classes		[]Class		`gorm:"ForeignKey:UserID" json:"classes,omitempty"`
	Forecast	Forecast	`gorm:"ForeignKey:UserID" json:"forecast,omitempty"`
}
func (e User) IsInitializable() bool { return false }

func (e User) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == os.Getenv("DEV_ENV") {
		db.DropTableIfExists(&Forecast{})
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