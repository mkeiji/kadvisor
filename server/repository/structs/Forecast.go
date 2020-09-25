package structs

import (
	"errors"
	"os"

	"gorm.io/gorm"
)

type Forecast struct {
	Base
	UserID  int             `json:"userID,omitempty"`
	Year    int             `json:"year,omitempty"`
	Entries []ForecastEntry `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"entries,omitempty"`
}

func (f Forecast) IsInitializable() bool { return false }

func (f Forecast) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == "DEV" {
		db.Migrator().DropTable(&ForecastEntry{})
		db.Migrator().DropTable(&Forecast{})
	}
	db.AutoMigrate(&Forecast{})
}

func (f Forecast) Initialize(db *gorm.DB) {}

/* GORM HOOKS */
func (f *Forecast) BeforeSave(db *gorm.DB) (err error) {
	err = f.validateEntriesMonth()
	if err == nil {
		err = f.isDuplicate(db)
	}
	if err == nil && f.Year == 0 {
		err = errors.New("year is required")
	}
	return
}

func (f *Forecast) BeforeDelete(db *gorm.DB) (err error) {
	err = db.First(&f, f.ID).Error
	return
}

/* PRIVATE */
func (f *Forecast) isDuplicate(db *gorm.DB) (err error) {
	var forecast Forecast
	fErr := db.Where(
		"user_id=? AND year=?",
		f.UserID,
		f.Year,
	).First(&forecast).Error
	if fErr == nil {
		err = errors.New("user already has a forecast")
	}
	return
}

func (f *Forecast) validateEntriesMonth() (err error) {
	var entriesMonth []int
	checked := map[int]bool{}

	for _, entry := range f.Entries {
		entriesMonth = append(entriesMonth, entry.Month)
	}

	for _, month := range entriesMonth {
		if checked[month] != true {
			checked[month] = true
		} else {
			err = errors.New("repeated month not allowed")
		}
	}

	return
}
