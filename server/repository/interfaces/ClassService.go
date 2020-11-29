package interfaces

import (
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
)

//go:generate mockgen -destination=mocks/mock_class_service.go -package=mocks . ClassService
type ClassService interface {
	GetClass(userID int, classID int) dtos.KhttpResponse
	GetManyByUserId(userID int) dtos.KhttpResponse
	GetOneById(id int) dtos.KhttpResponse
	Post(class s.Class) dtos.KhttpResponse
	Put(class s.Class) dtos.KhttpResponse
	Delete(id int) dtos.KhttpResponse
}
