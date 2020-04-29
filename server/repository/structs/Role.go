package structs

import (
	"github.com/jinzhu/gorm"
	"os"
)

type Role struct {
	Base
	RoleType	string    		`json:"roleType"`
	Description	string 			`json:"description"`
	Permission	[]Permission 	`gorm:"many2many:role_permissions;ForeignKey:ID" json:"permission"`
}

func (e Role) IsInitializable() bool { return true }

func (e Role) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == os.Getenv("DEV_ENV") {
		db.DropTableIfExists(&Permission{})
		db.DropTableIfExists(&Role{})
	}
	db.AutoMigrate(&Role{})
}

func (e Role) Initialize(db *gorm.DB) {
	role1 := Role{RoleType: "ADMIN"	, Description: "Admin"	, Permission: []Permission{{PermissionType: "VIEW"}, {PermissionType: "EDIT"}}}
	role2 := Role{RoleType: "REGULAR"	, Description: "Regular", Permission: []Permission{{PermissionType: "VIEW22"}}}
	db.Create(&role1)
	db.Create(&role2)
}

/* GORM HOOKS */
func (e *Role) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&e).Update("is_active", false)
	return
}
