package structs

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Login struct {
	Base
	UserID 		uint 		`json:"userID"`
	RoleID		uint 		`json:"roleID"`
	UserName	string		`json:"userName"`
	Password	string		`json:"passoword"`
	LastLogin	time.Time	`sql:"DEFAULT:current_timestamp"`
}
func (e Login) InitializeTable(db *gorm.DB) {
	db.Model(&Login{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Login{}).AddForeignKey("role_id", "roles(id)", "CASCADE", "CASCADE")
}
