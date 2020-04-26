package mappers

import (
	"github.com/jinzhu/gorm"
	"kadvisor/server/repository/structs"
)

// Mappers should not be calling the db.
// improve this in the future.
type UserMapper struct {}

func (u UserMapper) MapSubClassOnSave(
	user structs.User, db *gorm.DB) structs.User {

	id := int(user.ID)
	for i, class := range user.Classes {
		for j, subclass := range class.SubClasses {
			user.Classes[i].SubClasses[j].UserID = id
			db.Model(&subclass).Where(
				"id=?", subclass.ID).Update(
					"user_id", id)
		}
	}

	return user
}

func (u UserMapper) MapSubClassesOnLoad(
	users []structs.User, db *gorm.DB) []structs.User {
	var list []structs.User
	for _, user := range users {
		list = append(list, u.MapSubClassOnLoad(user, db))
	}
	return list
}

func (u UserMapper) MapSubClassOnLoad(
	user structs.User, db *gorm.DB) structs.User {
	for i, class := range user.Classes {
		var loadedClass structs.Class
		db.Preload("SubClasses").Find(&loadedClass, class.ID)
		user.Classes[i] = loadedClass
	}
	return user
}