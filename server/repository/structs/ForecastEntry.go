package structs

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type ForecastEntry struct {
	Base
	ForecastID 	int		`json:"forecastID,omitempty"`
	Month 		int		`json:"month,omitempty"`
	Income 		float64	`json:"income,omitempty"`
	Expense 	float64	`json:"expense,omitempty"`
}

func (f ForecastEntry) IsInitializable() bool {return false}

func (f ForecastEntry) Migrate(db *gorm.DB) {
	db.AutoMigrate(&ForecastEntry{})
}

func (f ForecastEntry) Initialize(db *gorm.DB) {}

/* GORM HOOKS */
func (f *ForecastEntry) BeforeCreate(db *gorm.DB) (err error) {
	err = f.validate(db)
	return
}

func (f *ForecastEntry) BeforeUpdate(db *gorm.DB) (err error) {
	err = f.validate(db)
	return
}

func (f *ForecastEntry) BeforeDelete(db *gorm.DB) (err error) {
	db.Model(&f).Update("is_active", false)
	return
}

/* PRIVATE */
func (f *ForecastEntry) validate(db *gorm.DB) (err error) {
	err = f.isDuplicate(db)
	if err == nil {
		err = f.validateMonth()
	}
	return
}

func (f *ForecastEntry) isDuplicate(db *gorm.DB) (err error) {
	var forecast ForecastEntry
	fErr := db.Where(
		"forecast_id=? AND month=? AND is_active=?",
		f.ForecastID,
		f.Month,
		true,
	).Find(&forecast).Error
	if fErr == nil {
		err = errors.New("month entry already exist")
	}
	return
}

func (f *ForecastEntry) validateMonth() (err error) {
	if f.Month < 1 || f.Month > 12 {
		err = errors.New("invalid month")
	}
	return
}