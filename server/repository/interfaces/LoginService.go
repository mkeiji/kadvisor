package interfaces

import (
	"kadvisor/server/libs/dtos"
	s "kadvisor/server/repository/structs"
)

//go:generate mockgen -destination=mocks/mock_login_service.go -package=mocks . LoginService
type LoginService interface {
	GetOneByEmail(email string) dtos.KhttpResponse
	Put(login s.Login) dtos.KhttpResponse
	UpdateLoginStatus(login s.Login, isLoggedIn bool) dtos.KhttpResponse
}
