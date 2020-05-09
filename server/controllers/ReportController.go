package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/services"
	"net/http"
	"strconv"
)

type ReportController struct {
	service services.ReportService
}

func (ctrl *ReportController) LoadEndpoints(router *gin.Engine) {
	// report types
	typeBalance := "BALANCE"
	typeYear 	:= "YTD"

	// get(/report?type)
	router.GET("/api/kadvisor/:uid/report", func (c *gin.Context) {
		userID		, _ := strconv.Atoi(c.Param("uid"))
		rType			:= c.Query("type")

		uErr := KeiUserUtil.ValidUser(userID)

		if uErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if rType == typeBalance {
			ctrl.getBalance(c, userID)
		} else if rType == typeYear {
			ctrl.getYearToDateReport(c, userID)
		}  else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "query param error"})
		}
	})
}

func (ctrl *ReportController) getBalance(c *gin.Context, userID int) {
	balance, err := ctrl.service.GetBalance(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, balance)
	}
}

func (ctrl *ReportController) getYearToDateReport(c *gin.Context, userID int) {
	yearly, err := ctrl.service.GetYearToDateReport(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, yearly)
	}
}