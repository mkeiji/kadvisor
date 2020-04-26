package structs

import "github.com/jinzhu/gorm"

type Permission struct {
	Base
	PermissionType	string	`json:"permissionType"`
	Description		string	`json:"description"`
}

func (e Permission) IsInitializable() bool { return false }

func (e Permission) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Permission{})
}

func (e Permission) Initialize(db *gorm.DB) {/* empty */}
