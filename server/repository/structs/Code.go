package structs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strings"
)

type Code struct {
	Base
	CodeTypeID string     `json:"codeType,omitempty"`
	CodeGroup  string     `json:"codeGroup,omitempty"`
	Name       string     `json:"name,omitempty"`
	CodeText   []CodeText `gorm:"ForeignKey:CodeID"`
}

func (e Code) IsInitializable() bool { return true }

func (e Code) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == "DEV" {
		db.DropTableIfExists(&CodeText{})
	}

	db.AutoMigrate(&Code{})
}

func (e Code) Initialize(db *gorm.DB) {
	e.insertCode(db, "INCOME_ENTRY_TYPE", "EntryTypeCodeID", "Income", "en")
	e.insertCode(db, "EXPENSE_ENTRY_TYPE", "EntryTypeCodeID", "Expense", "en")
}

/* GORM HOOKS */
func (e *Code) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&e).Update("is_active", false)
	return
}

func (e Code) insertCode(
	db *gorm.DB,
	codeTypeID string,
	codeGroup string,
	name string,
	locale string) {

	var code Code
	err := db.Where("code_type_id=?", codeTypeID).First(&code).Error
	if err != nil {
		code = e.createCode(codeTypeID, codeGroup, name, locale)
		db.Create(&code)
	}
}

func (e Code) createCode(
	codeTypeID string,
	codeGroup string,
	name string,
	locale string) Code {

	lowLoc := strings.ToLower(locale)
	textID := fmt.Sprintf("%v_%v", strings.ToUpper(lowLoc), codeTypeID)
	return Code{
		CodeTypeID: codeTypeID,
		CodeGroup:  codeGroup,
		Name:       name,
		CodeText: []CodeText{
			{
				TextID: textID,
				Locale: lowLoc,
			},
		},
	}
}
