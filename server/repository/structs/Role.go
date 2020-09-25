package structs

import (
	"gorm.io/gorm"
	"os"
	"strings"
)

type Role struct {
	Base
	RoleType    string       `json:"roleType,omitempty"`
	Description string       `json:"description,omitempty"`
	Permission  []Permission `gorm:"many2many:role_permissions;ForeignKey:ID" json:"permission,omitempty"`
}

func (e Role) IsInitializable() bool { return true }

func (e Role) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == "DEV" {
		db.Migrator().DropTable(&Permission{})
		db.Migrator().DropTable(&Role{})
	}
	db.AutoMigrate(&Role{})
}

func (e Role) Initialize(db *gorm.DB) {
	e.insertRole(db, "ADMIN", "Admin", []string{"VIEW", "EDIT"})
	e.insertRole(db, "REGULAR", "Regular", []string{"VIEW"})
}

func (e Role) insertRole(
	db *gorm.DB,
	roleType string,
	description string,
	permissions []string) {

	var role Role
	uppRoleType := strings.ToUpper(roleType)

	err := db.Where("role_type=?", uppRoleType).First(&role).Error
	if err != nil {
		role = e.createRole(uppRoleType, description, permissions)
		db.Create(&role)
	}
}

func (e Role) createRole(roleType string, description string, permissions []string) Role {
	return Role{
		RoleType:    strings.ToUpper(roleType),
		Description: description,
		Permission:  e.createPermissions(permissions),
	}
}

func (e Role) createPermissions(permissions []string) []Permission {
	var result []Permission
	for _, p := range permissions {
		result = append(result, Permission{PermissionType: strings.ToUpper(p)})
	}
	return result
}
