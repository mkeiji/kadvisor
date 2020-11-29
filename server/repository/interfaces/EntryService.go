package interfaces

import (
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
)

//go:generate mockgen -destination=mocks/mock_entry_service.go -package=mocks . EntryService
type EntryService interface {
	GetManyByUserId(userID int, limit int) dtos.KhttpResponse
	GetManyByClassId(classID int, limit int) dtos.KhttpResponse
	GetOneById(id int) dtos.KhttpResponse
	Post(entry s.Entry) dtos.KhttpResponse
	Put(entry s.Entry) dtos.KhttpResponse
	Delete(id int) dtos.KhttpResponse
}
