package interfaces

import s "kadvisor/server/repository/structs"

//go:generate mockgen -destination=mocks/mock_user_repository.go -package=mocks . UserRepository
type UserRepository interface {
	FindAll(preloaded bool) ([]s.User, error)
	FindOne(id int, preloaded bool) (s.User, error)
	Create(user s.User) (s.User, error)
	Update(user s.User) (s.User, error)
	Delete(userID int) (s.User, error)
}
