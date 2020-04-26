package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func (t *UserController) LoadEndpoints(router *gin.Engine) {
	// getOne(/user)
	router.GET("/api/user/:id", func (context *gin.Context) {
		isPreloaded, err := strconv.ParseBool(
			context.DefaultQuery("preloaded", "false"))

		userID, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		storedUser, err := t.userService.GetOne(userID, isPreloaded)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"user": storedUser})
		}
	})
	
	// getMany(/users)
	router.GET("/api/users", func (context *gin.Context) {
		isPreloaded, err := strconv.ParseBool(
			context.DefaultQuery("preloaded", "false"))

		users, err := t.userService.GetMany(isPreloaded)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		} else {
			context.JSON(http.StatusOK, gin.H{"users": users})
		}
	})

	// post(/user)
	router.POST("/api/user", func (context *gin.Context) {
		var user structs.User
		context.BindJSON(&user)
		KeiPassUtil.HashAndSalt(&user)
		savedUser, err := t.userService.Post(user)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		} else {
			context.JSON(http.StatusOK, gin.H{"user": savedUser})
		}
	})
}