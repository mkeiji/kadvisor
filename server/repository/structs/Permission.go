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

/* GORM HOOKS */
func (e *Permission) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&e).Update("is_active", false)
	return
}