package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"net/http"
	"strconv"
)

type ReportController struct {
	service services.ReportService
	auth    services.KeiAuthService
}

func (ctrl *ReportController) LoadEndpoints(router *gin.Engine) {
	// report types
	typeBalance := "BALANCE"
	typeYear := "YTD"
	typeYearFC := "YFC"

	reportRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	reportRoutes.Use(jwt.MiddlewareFunc())
	{
		// get(/report?type?year)
		reportRoutes.GET("/report", func(c *gin.Context) {
			userID, _ := strconv.Atoi(c.Param("uid"))
			year, _ := strconv.Atoi(c.Query("year"))
			rType := c.Query("type")

			uErr := KeiUserUtil.ValidUser(userID)

			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if rType == typeBalance {
				ctrl.getBalance(c, userID)
			} else if rType == typeYear && c.Query("year") != "" {
				ctrl.getYearToDateReport(c, userID, year)
			} else if rType == typeYearFC && c.Query("year") != "" {
				ctrl.getYearToDateWithForecast(c, userID, year)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "query param error"})
			}
		})

		// get(/reportavailable)
		reportRoutes.GET("/reportavailable", func(c *gin.Context) {
			userID, _ := strconv.Atoi(c.Param("uid"))
			uErr := KeiUserUtil.ValidUser(userID)

			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else {
				rAvailable, err := ctrl.service.GetReportAvailable(userID)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				} else {
					c.JSON(http.StatusOK, rAvailable)
				}
			}
		})
	}
}

func (ctrl *ReportController) getBalance(c *gin.Context, userID int) {
	balance, err := ctrl.service.GetBalance(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, balance)
	}
}

func (ctrl *ReportController) getYearToDateReport(c *gin.Context, userID int, year int) {
	yearly, err := ctrl.service.GetYearToDateReport(userID, year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, yearly)
	}
}

func (ctrl *ReportController) getYearToDateWithForecast(
	c *gin.Context, userID int, year int) {
	ytdFC, errors := ctrl.service.GetYearToDateWithForecastReport(userID, year)
	if len(errors) > 0 {
		errMsgs := ctrl.getErrorMessages(errors)
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsgs})
	} else {
		c.JSON(http.StatusOK, ytdFC)
	}
}

func (ctrl *ReportController) getErrorMessages(errors []error) []string {
	var errMsg []string
	for _, err := range errors {
		errMsg = append(errMsg, err.Error())
	}
	return errMsg
}
