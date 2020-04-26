package structs

import "github.com/jinzhu/gorm"

type SubClass struct {
	Base
	ClassID 	int		`json:"classID"`
	UserID 		int		`json:"userID"`
	Name 		string	`json:"name"`
	Description	string	`json:"description"`
}
func (e SubClass) IsInitializable() bool {return false}

func (e SubClass) Migrate(db *gorm.DB) {
	db.AutoMigrate(&SubClass{})
}

func (e SubClass) Initialize(db *gorm.DB) {}