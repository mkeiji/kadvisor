package controllers

import (
	"github.com/gin-gonic/gin"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"
)

type ForecastController struct {
	FcService         i.ForecastService
	UsrService        i.UserService
	Auth              i.KeiAuthService
	ValidationService i.ValidationService
}

func NewForecastController() ForecastController {
	return ForecastController{
		FcService:         s.NewForecastService(),
		UsrService:        s.NewUserService(),
		Auth:              s.NewKeiAuthService(),
		ValidationService: s.NewValidationService(),
	}
}

func (this ForecastController) LoadEndpoints(router *gin.Engine) {
	this.FcService = s.NewForecastService()
	this.UsrService = s.NewUserService()

	forecastRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := this.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	forecastRoutes.Use(jwt.MiddlewareFunc())
	{
		forecastRoutes.GET("/forecast", this.GetOneForecast)
		forecastRoutes.POST("/forecast", this.PostForecast)
		forecastRoutes.DELETE("/forecast", this.DeleteForecast)
	}
}

// getOne(/forecast?year&preloaded)
func (this ForecastController) GetOneForecast(c *gin.Context) {
	var response dtos.KhttpResponse

	userID, _ := strconv.Atoi(c.Param("uid"))
	year, _ := strconv.Atoi(c.Query("year"))
	isPreloaded, _ := strconv.ParseBool(
		c.DefaultQuery("preloaded", "false"))

	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	response = this.FcService.GetOne(userID, year, isPreloaded)
	c.JSON(response.Status, response.Body)
	return
}

// post(/forecast)
func (this ForecastController) PostForecast(c *gin.Context) {
	var forecast structs.Forecast

	userID, _ := strconv.Atoi(c.Param("uid"))
	response := this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	c.ShouldBindJSON(&forecast)
	response = this.ValidationService.GetResponse(
		v.NewForecastValidator(),
		forecast,
	)
	if u.IsOKresponse(response.Status) {
		response = this.FcService.Post(forecast)
	}

	c.JSON(response.Status, response.Body)
	return
}

// delete(/forecast?id)
func (this ForecastController) DeleteForecast(c *gin.Context) {
	var response dtos.KhttpResponse

	forecastID, _ := strconv.Atoi(c.Query("id"))
	userID, _ := strconv.Atoi(c.Param("uid"))

	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	response = this.FcService.Delete(forecastID)
	c.JSON(response.Status, response.Body)
	return
}
