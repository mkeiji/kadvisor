package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	userService services.UserService
	auth        services.KeiAuthService
}

func (ctrl *UserController) LoadEndpoints(router *gin.Engine) {
	userRoutes := router.Group("/api")
	permission := enums.ADMIN
	jwt, err := ctrl.auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	userRoutes.Use(jwt.MiddlewareFunc())
	{
		// getOne(/user?preloaded)
		userRoutes.GET("/user/:id", func(context *gin.Context) {
			isPreloaded, err := strconv.ParseBool(
				context.DefaultQuery("preloaded", "false"))

			userID, err := strconv.Atoi(context.Param("id"))
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			storedUser, err := ctrl.userService.GetOne(userID, isPreloaded)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, storedUser)
			}
		})

		// getMany(/users?preloaded)
		userRoutes.GET("/users", func(context *gin.Context) {
			isPreloaded, err := strconv.ParseBool(
				context.DefaultQuery("preloaded", "false"))

			users, err := ctrl.userService.GetMany(isPreloaded)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, users)
			}
		})

		// put(/user)
		userRoutes.PUT("/user", func(context *gin.Context) {
			var user structs.User
			context.BindJSON(&user)

			updated, err := ctrl.userService.Put(user)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, updated)
			}
		})

		// delete(/user)
		userRoutes.DELETE("/user/:id", func(context *gin.Context) {
			userID, err := strconv.Atoi(context.Param("id"))
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}

			deletedUser, dErr := ctrl.userService.Delete(userID)
			if dErr != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": dErr.Error()})
			} else {
				context.JSON(http.StatusOK, deletedUser)
			}
		})
	}

	// post(/user)
	router.POST("/api/user", func(context *gin.Context) {
		var user structs.User
		context.BindJSON(&user)
		KeiPassUtil.HashAndSalt(&user)
		savedUser, err := ctrl.userService.Post(user)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, savedUser)
		}
	})
}
