package repository

import (
	s "kadvisor/server/repository/structs"
	app "kadvisor/server/resources/application"

	"gorm.io/gorm"
)

type LoginRepository struct {
	Db *gorm.DB
}

func NewLoginRepository() LoginRepository {
	return LoginRepository{
		Db: app.Db,
	}
}

func (this LoginRepository) FindOneByEmail(email string) (s.Login, error) {
	var login s.Login

	err := this.Db.Where("email=?", email).First(&login).Error
	if err != nil {
		return login, err
	}
	return login, nil
}

func (this LoginRepository) Update(
	login s.Login,
) (s.Login, error) {
	stored, err := this.findOne(login.ID)
	if err == nil {
		err = this.Db.Model(&stored).Updates(login).Error
	}

	return stored, err
}

func (this LoginRepository) UpdateLoginStatus(
	login s.Login,
	isLoggedIn bool,
) (s.Login, error) {
	storedLogin, fErr := this.FindOneByEmail(login.Email)
	if fErr != nil {
		return storedLogin, fErr
	} else {
		err := this.Db.
			Model(&storedLogin).
			Where("email=?", login.Email).
			Update("IsLoggedIn", isLoggedIn).Error
		if err != nil {
			return s.Login{}, err
		}
	}

	return storedLogin, nil
}

func (this LoginRepository) findOne(
	id int,
) (s.Login, error) {
	var storedLogin s.Login
	err := this.Db.Where("id=?", id).First(&storedLogin).Error
	return storedLogin, err
}
