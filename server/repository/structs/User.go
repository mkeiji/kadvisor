package structs

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type User struct {
	Base
	IsPremium bool       `sql:"DEFAULT:FALSE" json:"isPremium,omitempty"`
	FirstName string     `json:"firstName,omitempty"`
	LastName  string     `json:"lastName,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Address   string     `json:"address,omitempty"`
	Login     Login      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"login,omitempty"`
	Entries   []Entry    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"entries,omitempty"`
	Classes   []Class    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"classes,omitempty"`
	Forecast  []Forecast `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"forecast,omitempty"`
}

func (this User) IsInitializable() bool {
	if os.Getenv("APP_ENV") == "DEV" {
		return true
	} else {
		return false
	}
}

func (this User) Migrate(db *gorm.DB) {
	if os.Getenv("APP_ENV") == "DEV" {
		db.Migrator().DropTable(&Forecast{})
		db.Migrator().DropTable(&Class{})
		db.Migrator().DropTable(&Entry{})
		db.Migrator().DropTable(&Login{})
		db.Migrator().DropTable(&User{})
	}
	db.AutoMigrate(&User{})
}

func (this User) Initialize(db *gorm.DB) {
	user := this.buildTestUser()
	err := db.Save(&user).Error
	if err != nil {
		log.Println(err)
	}
}

func (this User) buildTestUser() User {
	bytePwd := []byte("jkl")
	pwd, _ := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)

	return User{
		FirstName: "kguest",
		LastName:  "guestuser",
		IsPremium: true,
		Phone:     "111-111-1111",
		Address:   "Test Address",
		Login: Login{
			RoleID:   2,
			Email:    "test@email.com",
			UserName: "guest",
			Password: string(pwd),
		},
		Classes: []Class{
			{
				Name:        "Food",
				Description: "food",
			},
			{
				Name:        "Objects",
				Description: "objects",
			},
			{
				Name:        "Work",
				Description: "work",
			},
		},
	}
}
