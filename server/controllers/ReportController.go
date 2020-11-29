package controllers

import (
	"errors"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReportController struct {
	Service    i.ReportService
	UsrService i.UserService
	Auth       i.KeiAuthService
}

func NewReportController() ReportController {
	return ReportController{
		Service:    s.NewReportService(),
		UsrService: s.NewUserService(),
		Auth:       s.NewKeiAuthService(),
	}
}

func (this ReportController) LoadEndpoints(router *gin.Engine) {
	this.Service = s.NewReportService()
	this.UsrService = s.NewUserService()

	reportRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := this.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	reportRoutes.Use(jwt.MiddlewareFunc())
	{
		reportRoutes.GET("/report", this.GetReport)
		reportRoutes.GET("/reportavailable", this.GetReportAvailable)
	}
}

// get(/report?type?year)
func (this ReportController) GetReport(c *gin.Context) {
	// report types
	typeBalance := "BALANCE"
	typeYear := "YTD"
	typeYearFC := "YFC"

	var response dtos.KhttpResponse
	userID, _ := strconv.Atoi(c.Param("uid"))
	year, _ := strconv.Atoi(c.Query("year"))
	rType := c.Query("type")

	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	if rType == typeBalance {
		response = this.Service.GetBalance(userID)
	} else if rType == typeYear && year != 0 {
		response = this.Service.GetYearToDateReport(userID, year)
	} else if rType == typeYearFC && year != 0 {
		response = this.Service.GetYearToDateWithForecastReport(userID, year)
	} else {
		response = dtos.NewBadKresponse(errors.New("query param error"))
	}

	c.JSON(response.Status, response.Body)
	return
}

// get(/reportavailable?forecast)
func (this ReportController) GetReportAvailable(c *gin.Context) {
	var response dtos.KhttpResponse
	isForecast, _ := strconv.ParseBool(c.Query("forecast"))

	userID, _ := strconv.Atoi(c.Param("uid"))
	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	if isForecast == true {
		response = this.Service.GetReportForecastAvailable(userID)
	} else {
		response = this.Service.GetReportAvailable(userID)
	}
	c.JSON(response.Status, response.Body)
	return
}
