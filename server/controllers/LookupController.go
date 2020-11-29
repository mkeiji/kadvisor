package controllers

import (
	"github.com/gin-gonic/gin"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"
)

type LookupController struct {
	Service    i.LookupService
	UsrService i.UserService
	Auth       i.KeiAuthService
}

func NewLookupController() LookupController {
	return LookupController{
		Service:    s.NewLookupService(),
		UsrService: s.NewUserService(),
		Auth:       s.NewKeiAuthService(),
	}
}

func (this LookupController) LoadEndpoints(router *gin.Engine) {
	this.Service = s.NewLookupService()
	this.UsrService = s.NewUserService()

	lookupRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := this.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	lookupRoutes.Use(jwt.MiddlewareFunc())
	{
		// get(/lookup?codeGroup)
		lookupRoutes.GET("/lookup", this.GetLookup)
	}
}

func (this LookupController) GetLookup(c *gin.Context) {
	var response dtos.KhttpResponse

	userID, _ := strconv.Atoi(c.Param("uid"))
	codeGroup := c.Query("codeGroup")

	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	response = this.Service.GetAllByCodeGroup(codeGroup)
	c.JSON(response.Status, response.Body)
	return
}
