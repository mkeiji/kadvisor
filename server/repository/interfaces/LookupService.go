package interfaces

import (
	"kadvisor/server/libs/dtos"
)

//go:generate mockgen -destination=mocks/mock_lookup_service.go -package=mocks . LookupService
type LookupService interface {
	GetAllByCodeGroup(codeGroup string) dtos.KhttpResponse
}
