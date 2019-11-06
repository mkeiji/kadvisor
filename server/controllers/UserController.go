package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func (t *UserController) LoadEndpoints(router *gin.Engine) {
	// getMany(/users)
	router.GET("/api/users", func (context *gin.Context) {
		users := t.userService.GetMany()
		context.JSON(http.StatusOK, gin.H{"users": users})
	})

	// post(/user)
	router.POST("/api/user", func (context *gin.Context) {
		var user structs.User
		context.BindJSON(&user)
		t.userService.Post(user)
		context.JSON(http.StatusOK, gin.H{"user": user})
	})
}