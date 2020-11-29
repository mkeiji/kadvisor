package controllers

import (
	"github.com/gin-gonic/gin"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"
)

type ForecastEntryController struct {
	Service    i.ForecastEntryService
	UsrService i.UserService
	Auth       i.KeiAuthService
}

func NewForecastEntryController() ForecastEntryController {
	return ForecastEntryController{
		Service:    s.NewForecastEntryService(),
		UsrService: s.NewUserService(),
		Auth:       s.NewKeiAuthService(),
	}
}

func (this ForecastEntryController) LoadEndpoints(router *gin.Engine) {
	this.Service = s.NewForecastEntryService()
	this.UsrService = s.NewUserService()

	forecastEntryRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := this.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	forecastEntryRoutes.Use(jwt.MiddlewareFunc())
	{
		forecastEntryRoutes.PUT("/forecastentry", this.PutForecastEntry)
	}
}

// put(/forecastentry)
func (this ForecastEntryController) PutForecastEntry(c *gin.Context) {
	var response dtos.KhttpResponse
	var entry structs.ForecastEntry

	userID, _ := strconv.Atoi(c.Param("uid"))
	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	c.ShouldBindJSON(&entry)
	response = this.Service.Put(entry)
	c.JSON(response.Status, response.Body)
	return
}
