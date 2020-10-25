package structs

import (
	"gorm.io/gorm"
	"os"
)

type User struct {
	Base
	IsPremium bool       `sql:"DEFAULT:FALSE" json:"isPremium,omitempty"`
	FirstName string     `json:"firstName,omitempty" validate:"required"`
	LastName  string     `json:"lastName,omitempty" validate:"required"`
	Phone     string     `json:"phone,omitempty"`
	Address   string     `json:"address,omitempty"`
	Login     Login      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"login,omitempty"`
	Entries   []Entry    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"entries,omitempty"`
	Classes   []Class    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"classes,omitempty"`
	Forecast  []Forecast `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"forecast,omitempty"`
}

func (e User) IsInitializable() bool { return false }

func (e User) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == "DEV" {
		db.Migrator().DropTable(&Forecast{})
		db.Migrator().DropTable(&Class{})
		db.Migrator().DropTable(&Entry{})
		db.Migrator().DropTable(&Login{})
		db.Migrator().DropTable(&User{})
	}
	db.AutoMigrate(&User{})
}

func (e User) Initialize(db *gorm.DB) { /* empty */ }
