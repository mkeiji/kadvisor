package interfaces

import (
	"kadvisor/server/libs/dtos"
)

//go:generate mockgen -destination=mocks/mock_validation_service.go -package=mocks . ValidationService
type ValidationService interface {
	GetResponse(validator Validator, obj interface{}) dtos.KhttpResponse
}
