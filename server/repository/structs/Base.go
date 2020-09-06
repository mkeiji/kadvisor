package structs

import "time"

type Base struct {
	ID        int        `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`
	IsActive  bool       `sql:"DEFAULT:TRUE" json:"isActive"`
}
