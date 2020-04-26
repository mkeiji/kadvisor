package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Login struct {
	Base
	UserID 		uint 		`json:"userID"`
	RoleID		uint 		`json:"roleID"`
	Email		string		`json:"email" gorm:"unique;not null"`
	UserName	string		`json:"userName"`
	Password	string		`json:"password"`
	IsLoggedIn	bool		`json:"isLoggedIn" sql:"DEFAULT:FALSE" `
	LastLogin	time.Time	`sql:"DEFAULT:current_timestamp"`
}

func (e Login) IsInitializable() bool { return false }

func (e Login) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Login{})
}

func (e Login) Initialize(db *gorm.DB) {/* empty */}
