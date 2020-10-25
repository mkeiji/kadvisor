package structs

import (
	"os"

	"gorm.io/gorm"
)

type Forecast struct {
	Base
	UserID  int             `json:"userID,omitempty"`
	Year    int             `json:"year,omitempty" validate:"required"`
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
