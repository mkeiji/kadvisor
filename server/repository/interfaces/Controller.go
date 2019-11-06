package interfaces

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	LoadEndpoints(router *gin.Engine)
}
