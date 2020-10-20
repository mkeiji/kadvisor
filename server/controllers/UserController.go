package controllers

import (
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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

	// post(/user) - unprotected by jwt
	router.POST("/api/user", func(context *gin.Context) {
		var user structs.User
		context.BindJSON(&user)
		KeiPassUtil.HashAndSalt(&user)
		response := ctrl.userService.Post(user)
		context.JSON(response.Status, response.Body)
		return
	})

	userRoutes.Use(jwt.MiddlewareFunc())
	{
		// getOne(/user?preloaded)
		userRoutes.GET("/user/:id", func(context *gin.Context) {
			isPreloaded, _ := strconv.ParseBool(
				context.DefaultQuery("preloaded", "false"),
			)
			userID, _ := strconv.Atoi(context.Param("id"))

			response := ctrl.userService.GetOne(userID, isPreloaded)
			context.JSON(response.Status, response.Body)
			return
		})

		// getMany(/users?preloaded)
		userRoutes.GET("/users", func(context *gin.Context) {
			isPreloaded, _ := strconv.ParseBool(
				context.DefaultQuery("preloaded", "false"),
			)

			response := ctrl.userService.GetMany(isPreloaded)
			context.JSON(response.Status, response.Body)
			return
		})

		// put(/user)
		userRoutes.PUT("/user", func(context *gin.Context) {
			var user structs.User
			context.BindJSON(&user)

			response := ctrl.userService.Put(user)
			context.JSON(response.Status, response.Body)
			return
		})

		// delete(/user)
		userRoutes.DELETE("/user/:id", func(context *gin.Context) {
			userID, _ := strconv.Atoi(context.Param("id"))

			response := ctrl.userService.Delete(userID)
			context.JSON(response.Status, response.Body)
			return
		})
	}
}
