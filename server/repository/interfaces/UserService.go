package interfaces

import (
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
)

//go:generate mockgen -destination=mocks/mock_user_service.go -package=mocks . UserService
type UserService interface {
	GetMany(preloaded bool) dtos.KhttpResponse
	GetOne(id int, preloaded bool) dtos.KhttpResponse
	Post(user s.User) dtos.KhttpResponse
	Put(user s.User) dtos.KhttpResponse
	Delete(userID int) dtos.KhttpResponse
}
