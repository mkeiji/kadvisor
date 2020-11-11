package controllers

import (
	"github.com/gin-gonic/gin"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"
)

type ForecastEntryController struct {
	Service    s.ForecastEntryService
	UsrService s.UserService
	Auth       s.KeiAuthService
}

func (ctrl ForecastEntryController) LoadEndpoints(router *gin.Engine) {
	ctrl.Service = s.NewForecastEntryService()
	ctrl.UsrService = s.NewUserService()

	forecastEntryRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.Auth.GetAuthUtil(permission)
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
			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&entry)
			response = ctrl.Service.Put(entry)
			c.JSON(response.Status, response.Body)
			return
		})
	}
}
