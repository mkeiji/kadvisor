package interfaces

import s "kadvisor/server/repository/structs"

//go:generate mockgen -destination=mocks/mock_code_code_text_repository.go -package=mocks . CodeCodeTextRepository
type CodeCodeTextRepository interface {
	FindAllByCodeGroup(codeGroup string) ([]s.Code, error)
	FindOne(codeTypeID string) (s.Code, error)
}
