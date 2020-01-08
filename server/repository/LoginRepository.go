package repository

import (
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/application"
)

type LoginRepository struct {}

func (l *LoginRepository) FindOneByEmail(email string) structs.Login {
	var login structs.Login
	application.Db.Where("email=?", email).First(&login)
	return login
}