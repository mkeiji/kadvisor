package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"net/http"
	"strconv"
)

type ForecastController struct {
	service services.ForecastService
	auth    services.KeiAuthService
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
			userID, _ := strconv.Atoi(c.Param("uid"))
			year, _ := strconv.Atoi(c.Query("year"))
			isPreloaded, _ := strconv.ParseBool(
				c.DefaultQuery("preloaded", "false"))

			uErr := KeiUserUtil.ValidUser(userID)

			forecast, fErr := ctrl.service.GetOne(userID, year, isPreloaded)
			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if year == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "year param is required"})
			} else if fErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": fErr.Error()})
			} else {
				c.JSON(http.StatusOK, forecast)
			}
		})

		// post(/forecast)
		forecastRoutes.POST("/forecast", func(c *gin.Context) {
			var forecast structs.Forecast
			c.BindJSON(&forecast)

			saveForecast, err := ctrl.service.Post(forecast)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, saveForecast)
			}
		})

		// delete(/forecast?id)
		forecastRoutes.DELETE("/forecast", func(c *gin.Context) {
			forecastID, _ := strconv.Atoi(c.Query("id"))
			userID, _ := strconv.Atoi(c.Param("uid"))

			uErr := KeiUserUtil.ValidUser(userID)

			deleted, err := ctrl.service.Delete(forecastID)
			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else if forecastID == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("missing id param")})
			} else {
				c.JSON(http.StatusOK, deleted)
			}
		})
	}
}
