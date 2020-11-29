package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	i "kadvisor/server/repository/interfaces"
	s "kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	svc "kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassController struct {
	Service           i.ClassService
	Auth              i.KeiAuthService
	UsrService        i.UserService
	ValidationService i.ValidationService
}

func NewClassController() ClassController {
	return ClassController{
		Service:           svc.NewClassService(),
		Auth:              svc.NewKeiAuthService(),
		UsrService:        svc.NewUserService(),
		ValidationService: svc.NewValidationService(),
	}
}

func (this ClassController) LoadEndpoints(router *gin.Engine) {
	this.Service = svc.NewClassService()
	this.UsrService = svc.NewUserService()

	classRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := this.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	classRoutes.Use(jwt.MiddlewareFunc())
	{
		classRoutes.GET("/class", this.GetOneById)
		classRoutes.POST("/class", this.PostClass)
		classRoutes.PUT("/class", this.PutClass)
		classRoutes.DELETE("/class", this.DeleteClass)
	}
}

// getOne(/class?id)
func (this ClassController) GetOneById(c *gin.Context) {
	var response dtos.KhttpResponse
	userID, _ := strconv.Atoi(c.Param("uid"))
	classID, _ := strconv.Atoi(c.Query("id"))

	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	response = this.Service.GetClass(userID, classID)
	c.JSON(response.Status, response.Body)
	return
}

// post(/class)
func (this ClassController) PostClass(c *gin.Context) {
	var response dtos.KhttpResponse
	var class s.Class

	userID, _ := strconv.Atoi(c.Param("uid"))
	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	if bindErr := c.ShouldBindJSON(&class); bindErr != nil {
		response = dtos.NewBadKresponse(bindErr)
		return
	}

	response = this.ValidationService.GetResponse(
		v.NewClassValidator(),
		class,
	)
	if u.IsOKresponse(response.Status) {
		response = this.Service.Post(class)
	}

	c.JSON(response.Status, response.Body)
	return
}

// put(/class)
func (this ClassController) PutClass(c *gin.Context) {
	var response dtos.KhttpResponse
	var class s.Class

	userID, _ := strconv.Atoi(c.Param("uid"))
	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	if bindErr := c.ShouldBindJSON(&class); bindErr != nil {
		response = dtos.NewBadKresponse(bindErr)
		return
	}

	response = this.ValidationService.GetResponse(
		v.NewClassValidator(),
		class,
	)
	if u.IsOKresponse(response.Status) {
		response = this.Service.Put(class)
	}

	c.JSON(response.Status, response.Body)
	return
}

// delete(/class?id)
func (this ClassController) DeleteClass(c *gin.Context) {
	var response dtos.KhttpResponse

	classID, _ := strconv.Atoi(c.Query("id"))
	userID, _ := strconv.Atoi(c.Param("uid"))

	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	response = this.Service.Delete(classID)
	c.JSON(response.Status, response.Body)
	return
}
