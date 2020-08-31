package structs

import (
	"errors"
	"github.com/jinzhu/gorm"
	"os"
)

type Forecast struct {
	Base
	UserID  int             `json:"userID,omitempty"`
	Year    int             `json:"year,omitempty"`
	Entries []ForecastEntry `gorm:"ForeignKey:ForecastID" json:"entries,omitempty"`
}

func (f Forecast) IsInitializable() bool { return false }

func (f Forecast) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == os.Getenv("DEV_ENV") {
		db.DropTableIfExists(&ForecastEntry{})
		db.DropTableIfExists(&Forecast{})
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
	err = db.Find(&f, f.ID).Error
	if err == nil {
		db.Model(&f).Update("is_active", false)
		err = f.deleteChildren(db)
	}
	return
}

/* PRIVATE */
func (f *Forecast) isDuplicate(db *gorm.DB) (err error) {
	var forecast Forecast
	fErr := db.Where(
		"user_id=? AND year=? AND is_active=?", f.UserID, f.Year, true).Find(
		&forecast).Error
	if fErr == nil {
		err = errors.New("user already has a forecast")
	}
	return
}

func (f *Forecast) deleteChildren(db *gorm.DB) (err error) {
	return db.Where("forecast_id=?", f.ID).Delete(ForecastEntry{}).Error
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
