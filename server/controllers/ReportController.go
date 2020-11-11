package controllers

import (
	"errors"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	Service    s.ReportService
	UsrService s.UserService
	Auth       s.KeiAuthService
}

func (ctrl ReportController) LoadEndpoints(router *gin.Engine) {
	ctrl.Service = s.NewReportService()
	ctrl.UsrService = s.NewUserService()

	// report types
	typeBalance := "BALANCE"
	typeYear := "YTD"
	typeYearFC := "YFC"

	reportRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	reportRoutes.Use(jwt.MiddlewareFunc())
	{
		// get(/report?type?year)
		reportRoutes.GET("/report", func(c *gin.Context) {
			var response dtos.KhttpResponse
			userID, _ := strconv.Atoi(c.Param("uid"))
			year, _ := strconv.Atoi(c.Query("year"))
			rType := c.Query("type")

			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			if rType == typeBalance {
				response = ctrl.Service.GetBalance(userID)
			} else if rType == typeYear && year != 0 {
				response = ctrl.Service.GetYearToDateReport(userID, year)
			} else if rType == typeYearFC && year != 0 {
				response = ctrl.Service.GetYearToDateWithForecastReport(userID, year)
			} else {
				response = dtos.NewBadKresponse(errors.New("query param error"))
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// get(/reportavailable?forecast)
		reportRoutes.GET("/reportavailable", func(c *gin.Context) {
			var response dtos.KhttpResponse
			isForecast, _ := strconv.ParseBool(c.Query("forecast"))

			userID, _ := strconv.Atoi(c.Param("uid"))
			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			if isForecast == true {
				response = ctrl.Service.GetReportForecastAvailable(userID)
			} else {
				response = ctrl.Service.GetReportAvailable(userID)
			}
			c.JSON(response.Status, response.Body)
			return
		})
	}
}
