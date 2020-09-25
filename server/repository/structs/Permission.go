package structs

import "gorm.io/gorm"

type Permission struct {
	Base
	PermissionType string `json:"permissionType,omitempty"`
	Description    string `json:"description,omitempty"`
}

func (e Permission) IsInitializable() bool { return false }

func (e Permission) Migrate(db *gorm.DB) {
	db.AutoMigrate(&Permission{})
}

func (e Permission) Initialize(db *gorm.DB) { /* empty */ }
