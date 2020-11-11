package interfaces

import s "kadvisor/server/repository/structs"

//go:generate mockgen -destination=mocks/mock_entry_repository.go -package=mocks . EntryRepository
type EntryRepository interface {
	FindAllByUserId(userID int, limit int) ([]s.Entry, error)
	FindAllByClassId(classID int, limit int) ([]s.Entry, error)
	FindOne(id int) (s.Entry, error)
	Create(entry s.Entry) (s.Entry, error)
	Update(entry s.Entry) (s.Entry, error)
	Delete(id int) (int, error)
}
