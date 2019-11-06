package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller interface {
	LoadEndpoints(router *gin.Engine, db *gorm.DB)
}
