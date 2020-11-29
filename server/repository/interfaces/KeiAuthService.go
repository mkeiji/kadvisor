package interfaces

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"kadvisor/server/resources/enums"
)

//go:generate mockgen -destination=mocks/mock_kei_auth_service.go -package=mocks . KeiAuthService
type KeiAuthService interface {
	GetAuthUtil(permissionLevel enums.RoleEnum) (*jwt.GinJWTMiddleware, error)
}
