package interfaces

import s "kadvisor/server/repository/structs"

//go:generate mockgen -destination=mocks/mock_class_repository.go -package=mocks . ClassRepository
type ClassRepository interface {
	FindAllByUserId(userID int) ([]s.Class, error)
	FindOne(classID int) (s.Class, error)
	Create(class s.Class) (s.Class, error)
	Update(class s.Class) (s.Class, error)
	Delete(classID int) (int, error)
}
