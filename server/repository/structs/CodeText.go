package structs

import "gorm.io/gorm"

type CodeText struct {
	Base
	TextID string `json:"codeTextID,omitempty"`
	CodeID int    `json:"codeID,omitempty"`
	Locale string `json:"locale,omitempty"`
}

func (e CodeText) IsInitializable() bool { return false }

func (e CodeText) Migrate(db *gorm.DB) {
	db.AutoMigrate(&CodeText{})
}

func (e CodeText) Initialize(db *gorm.DB) {}
