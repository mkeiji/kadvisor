package controllers

import (
	"github.com/gin-gonic/gin"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	"kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"strconv"
)

type ForecastController struct {
	fcService         services.ForecastService
	usrService        services.UserService
	auth              services.KeiAuthService
	validator         validators.ForecastValidator
	validationService services.ValidationService
}

func (ctrl *ForecastController) LoadEndpoints(router *gin.Engine) {
	forecastRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	forecastRoutes.Use(jwt.MiddlewareFunc())
	{
		// getOne(/forecast?year&preloaded)
		forecastRoutes.GET("/forecast", func(c *gin.Context) {
			var response dtos.KhttpResponse

			userID, _ := strconv.Atoi(c.Param("uid"))
			year, _ := strconv.Atoi(c.Query("year"))
			isPreloaded, _ := strconv.ParseBool(
				c.DefaultQuery("preloaded", "false"))

			response = ctrl.usrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			response = ctrl.fcService.GetOne(userID, year, isPreloaded)
			c.JSON(response.Status, response.Body)
			return
		})

		// post(/forecast)
		forecastRoutes.POST("/forecast", func(c *gin.Context) {
			var forecast structs.Forecast

			userID, _ := strconv.Atoi(c.Param("uid"))
			response := ctrl.usrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&forecast)
			response = ctrl.validationService.GetResponse(
				ctrl.validator,
				forecast,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.fcService.Post(forecast)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// delete(/forecast?id)
		forecastRoutes.DELETE("/forecast", func(c *gin.Context) {
			var response dtos.KhttpResponse

			forecastID, _ := strconv.Atoi(c.Query("id"))
			userID, _ := strconv.Atoi(c.Param("uid"))

			response = ctrl.usrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			response = ctrl.fcService.Delete(forecastID)
			c.JSON(response.Status, response.Body)
			return
		})
	}
}
