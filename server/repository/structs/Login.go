package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Login struct {
	Base
	UserID 		int 		`json:"userID,omitempty"`
	RoleID		int 		`json:"roleID,omitempty"`
	Email		string		`json:"email,omitempty" gorm:"unique;not null"`
	UserName	string		`json:"userName,omitempty"`
	Password	string		`json:"password,omitempty"`
	IsLoggedIn	bool		`json:"isLoggedIn,omitempty"`
	LastLogin	*time.Time	`json:"lastLogin,omitempty"`
}

func (e Login) IsInitializable() bool { return false }

func (e Login) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Login{})
}

func (e Login) Initialize(db *gorm.DB) {/* empty */}

/* GORM HOOKS */
func (e *Login) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&e).Update("is_active", false)
	return
}