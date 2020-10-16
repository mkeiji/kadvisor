package structs

import (
	"errors"

	"gorm.io/gorm"
)

type ForecastEntry struct {
	Base
	ForecastID int     `json:"forecastID,omitempty"`
	Month      int     `json:"month,omitempty"`
	Income     float64 `json:"income,omitempty"`
	Expense    float64 `json:"expense,omitempty"`
}

func (f ForecastEntry) IsInitializable() bool { return false }

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
	err = f.exists(db)
	if err != nil {
		return
	}
	err = f.validateMonth()
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

func (f *ForecastEntry) exists(db *gorm.DB) (err error) {
	var forecast ForecastEntry
	err = db.Where("forecast_id=?", f.ForecastID).First(&forecast).Error
	return
}

func (f *ForecastEntry) isDuplicate(db *gorm.DB) (err error) {
	var forecast ForecastEntry
	fErr := db.Where(
		"forecast_id=? AND month=?",
		f.ForecastID,
		f.Month,
	).First(&forecast).Error
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
