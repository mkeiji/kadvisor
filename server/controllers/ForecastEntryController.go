package controllers

import (
	"github.com/gin-gonic/gin"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"strconv"
)

type ForecastEntryController struct {
	service    services.ForecastEntryService
	usrService services.UserService
	auth       services.KeiAuthService
}

func (ctrl *ForecastEntryController) LoadEndpoints(router *gin.Engine) {
	forecastEntryRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	forecastEntryRoutes.Use(jwt.MiddlewareFunc())
	{
		// put(/forecastentry)
		forecastEntryRoutes.PUT("/forecastentry", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var entry structs.ForecastEntry

			userID, _ := strconv.Atoi(c.Param("uid"))
			response = ctrl.usrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&entry)
			response = ctrl.service.Put(entry)
			c.JSON(response.Status, response.Body)
			return
		})
	}
}
