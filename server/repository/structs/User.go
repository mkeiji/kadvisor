package structs

import "github.com/jinzhu/gorm"

type User struct {
	Base
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Login     Login  `gorm:"ForeignKey:UserID" json:"login"`
}
func (e User) InitializeTable(db *gorm.DB) {
	//empty
}