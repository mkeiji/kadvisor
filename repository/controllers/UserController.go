package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"kadvisor/repository/structs"
)

type UserController struct {}

func (uc *UserController) LoadEndpoints(router *gin.Engine, db *gorm.DB) {
	// getMany(/users)
	router.GET("/users", func (context *gin.Context) {
		var users []structs.User
		if err := db.Find(&users).Error; err != nil {
			context.AbortWithStatus(404)
			fmt.Println(err)
		} else {
			context.JSON(200, users)
		}
	})

	// post(/user)
	router.POST("/user", func (context *gin.Context) {
		var user structs.User
		context.BindJSON(&user)
		db.Create(&user)
		context.JSON(200, user)
	})
}