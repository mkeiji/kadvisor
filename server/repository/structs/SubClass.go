package structs

import "github.com/jinzhu/gorm"

type SubClass struct {
	Base
	ClassID 	int		`json:"classID,omitempty"`
	UserID 		int		`json:"userID,omitempty"`
	Name 		string	`json:"name,omitempty"`
	Description	string	`json:"description,omitempty"`
}
func (e SubClass) IsInitializable() bool {return false}

func (e SubClass) Migrate(db *gorm.DB) {
	db.AutoMigrate(&SubClass{})
}

func (e SubClass) Initialize(db *gorm.DB) {}

/* GORM HOOKS */
func (e *SubClass) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&e).Update("is_active", false)
	return
}