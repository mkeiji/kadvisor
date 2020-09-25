package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
	"strconv"
)

type UserController struct {
	userService services.UserService
}

func (t *UserController) LoadEndpoints(router *gin.Engine) {
	// getOne(/user?preloaded)
	router.GET("/api/user/:id", func(context *gin.Context) {
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
			context.JSON(http.StatusOK, storedUser)
		}
	})

	// getMany(/users?preloaded)
	router.GET("/api/users", func(context *gin.Context) {
		isPreloaded, err := strconv.ParseBool(
			context.DefaultQuery("preloaded", "false"))

		users, err := t.userService.GetMany(isPreloaded)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, users)
		}
	})

	// post(/user)
	router.POST("/api/user", func(context *gin.Context) {
		var user structs.User
		context.BindJSON(&user)
		KeiPassUtil.HashAndSalt(&user)
		savedUser, err := t.userService.Post(user)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, savedUser)
		}
	})

	// put(/user)
	router.PUT("/api/user", func(context *gin.Context) {
		var user structs.User
		context.BindJSON(&user)

		updated, err := t.userService.Put(user)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, updated)
		}
	})

	// delete(/user)
	router.DELETE("/api/user/:id", func(context *gin.Context) {
		userID, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		deletedUser, dErr := t.userService.Delete(userID)
		if dErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": dErr.Error()})
		} else {
			context.JSON(http.StatusOK, deletedUser)
		}
	})
}
