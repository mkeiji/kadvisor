package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/libs/dtos"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService       i.UserService
	Auth              i.KeiAuthService
	ValidationService i.ValidationService
}

func NewUserController() UserController {
	return UserController{
		UserService:       s.NewUserService(),
		Auth:              s.NewKeiAuthService(),
		ValidationService: s.NewValidationService(),
	}
}

func (this UserController) LoadEndpoints(router *gin.Engine) {
	this.UserService = s.NewUserService()

	userRoutes := router.Group("/api")
	minPermission := enums.ADMIN
	jwt, err := this.Auth.GetAuthUtil(minPermission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	// unprotected by jwt
	router.POST("/api/user", this.PostUser)

	userRoutes.Use(jwt.MiddlewareFunc())
	{
		userRoutes.GET("/user/:id", this.GetOneUser)
		userRoutes.GET("/users", this.GetManyUsers)
		userRoutes.PUT("/user", this.UpdateUser)
		userRoutes.DELETE("/user/:id", this.DeleteUser)
	}
}

// postUser(/user)
func (ctrl UserController) PostUser(c *gin.Context) {
	var response dtos.KhttpResponse
	var user structs.User

	c.BindJSON(&user)
	response = ctrl.ValidationService.GetResponse(
		v.NewUserValidator(),
		user,
	)
	if u.IsOKresponse(response.Status) {
		hashedPwd, _ := KeiPassUtil.HashAndSalt(&user)
		user.Login.Password = hashedPwd
		response = ctrl.UserService.Post(user)
	}

	c.JSON(response.Status, response.Body)
	return
}

// getOne(/user?preloaded)
func (this UserController) GetOneUser(context *gin.Context) {
	isPreloaded, _ := strconv.ParseBool(
		context.DefaultQuery("preloaded", "false"),
	)
	userID, _ := strconv.Atoi(context.Param("id"))

	response := this.UserService.GetOne(userID, isPreloaded)
	context.JSON(response.Status, response.Body)
	return
}

// getMany(/users?preloaded)
func (this UserController) GetManyUsers(context *gin.Context) {
	isPreloaded, _ := strconv.ParseBool(
		context.DefaultQuery("preloaded", "false"),
	)

	response := this.UserService.GetMany(isPreloaded)
	context.JSON(response.Status, response.Body)
	return
}

// put(/user)
func (this UserController) UpdateUser(context *gin.Context) {
	var response dtos.KhttpResponse
	var user structs.User

	context.BindJSON(&user)
	response = this.ValidationService.GetResponse(
		v.NewUserValidator(),
		user,
	)
	if u.IsOKresponse(response.Status) {
		response = this.UserService.Put(user)
	}

	context.JSON(response.Status, response.Body)
	return
}

// delete(/user)
func (this UserController) DeleteUser(context *gin.Context) {
	userID, _ := strconv.Atoi(context.Param("id"))

	response := this.UserService.Delete(userID)
	context.JSON(response.Status, response.Body)
	return
}
