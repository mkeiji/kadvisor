package interfaces

import s "kadvisor/server/repository/structs"

//go:generate mockgen -destination=mocks/mock_login_repository.go -package=mocks . LoginRepository
type LoginRepository interface {
	FindOneByEmail(email string) (s.Login, error)
	Update(login s.Login) (s.Login, error)
	UpdateLoginStatus(login s.Login, isLoggedIn bool) (s.Login, error)
}
