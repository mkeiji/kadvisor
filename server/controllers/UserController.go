package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService       s.UserService
	Auth              s.KeiAuthService
	ValidationService s.ValidationService
}

func (ctrl UserController) LoadEndpoints(router *gin.Engine) {
	ctrl.UserService = s.NewUserService()

	userRoutes := router.Group("/api")
	permission := enums.ADMIN
	jwt, err := ctrl.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	// post(/user) - unprotected by jwt
	router.POST("/api/user", func(context *gin.Context) {
		var response dtos.KhttpResponse
		var user structs.User

		context.BindJSON(&user)
		response = ctrl.ValidationService.GetResponse(
			v.NewUserValidator(),
			user,
		)
		if u.IsOKresponse(response.Status) {
			hashedPwd, _ := KeiPassUtil.HashAndSalt(&user)
			user.Login.Password = hashedPwd
			response = ctrl.UserService.Post(user)
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

			response := ctrl.UserService.GetOne(userID, isPreloaded)
			context.JSON(response.Status, response.Body)
			return
		})

		// getMany(/users?preloaded)
		userRoutes.GET("/users", func(context *gin.Context) {
			isPreloaded, _ := strconv.ParseBool(
				context.DefaultQuery("preloaded", "false"),
			)

			response := ctrl.UserService.GetMany(isPreloaded)
			context.JSON(response.Status, response.Body)
			return
		})

		// put(/user)
		userRoutes.PUT("/user", func(context *gin.Context) {
			var response dtos.KhttpResponse
			var user structs.User

			context.BindJSON(&user)
			response = ctrl.ValidationService.GetResponse(
				v.NewUserValidator(),
				user,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.UserService.Put(user)
			}

			context.JSON(response.Status, response.Body)
			return
		})

		// delete(/user)
		userRoutes.DELETE("/user/:id", func(context *gin.Context) {
			userID, _ := strconv.Atoi(context.Param("id"))

			response := ctrl.UserService.Delete(userID)
			context.JSON(response.Status, response.Body)
			return
		})
	}
}
