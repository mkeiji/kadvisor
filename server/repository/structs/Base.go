package structs

type Base struct {
	ID        int   `json:"id" gorm:"primary_key"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}
