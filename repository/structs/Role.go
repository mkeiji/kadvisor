package structs

import "github.com/jinzhu/gorm"

type Role struct {
	Base
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Login		[]Login `gorm:"ForeignKey:RoleID" json:"login"`
}

func (e Role) InitializeTable(db *gorm.DB) {
	addUserRoles(db)
}

func addUserRoles(db *gorm.DB) {
	role1 := Role{Name: "Admin", Description: "Admin"}
	role2 := Role{Name: "Regular", Description: "Regular"}
	db.Create(&role1)
	db.Create(&role2)
}
