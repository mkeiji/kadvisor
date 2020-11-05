package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService       services.UserService
	auth              services.KeiAuthService
	validationService services.ValidationService
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
		var response dtos.KhttpResponse
		var user structs.User

		context.BindJSON(&user)
		response = ctrl.validationService.GetResponse(
			v.NewUserValidator(),
			user,
		)
		if u.IsOKresponse(response.Status) {
			KeiPassUtil.HashAndSalt(&user)
			response = ctrl.userService.Post(user)
		}

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
			var response dtos.KhttpResponse
			var user structs.User

			context.BindJSON(&user)
			response = ctrl.validationService.GetResponse(
				v.NewUserValidator(),
				user,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.userService.Put(user)
			}

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
