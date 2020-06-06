package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
	"strconv"
)

type ForecastController struct {
	service 	services.ForecastService
}

func (ctrl *ForecastController) LoadEndpoints(router *gin.Engine) {
	// getOne(/forecast?year&preloaded)
	router.GET("/api/kadvisor/:uid/forecast", func (c *gin.Context) {
		userID		, _	:= strconv.Atoi(c.Param("uid"))
		year		, _	:= strconv.Atoi(c.Query("year"))
		isPreloaded	, _	:= strconv.ParseBool(
			c.DefaultQuery("preloaded", "false"))

		uErr := KeiUserUtil.ValidUser(userID)

		forecast, fErr := ctrl.service.GetOne(userID, year, isPreloaded)
		if uErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if year == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "year param is required"})
		}  else if fErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fErr.Error()})
		} else {
			c.JSON(http.StatusOK, forecast)
		}
	})

	// post(/forecast)
	router.POST("/api/kadvisor/:uid/forecast", func (c *gin.Context) {
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
	router.DELETE("/api/kadvisor/:uid/forecast", func (c *gin.Context) {
		forecastID	, _ := strconv.Atoi(c.Query("id"))
		userID		, _ := strconv.Atoi(c.Param("uid"))

		uErr := KeiUserUtil.ValidUser(userID)

		deleted, err := ctrl.service.Delete(forecastID)
		if uErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if forecastID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("missing id param")})
		}else {
			c.JSON(http.StatusOK, deleted)
		}
	})
}
