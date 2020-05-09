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
	// get(/report?type)
	router.GET("/api/kadvisor/:uid/report", func (c *gin.Context) {
		userID		, _ := strconv.Atoi(c.Param("uid"))
		rType			:= c.Query("type")

		uErr := KeiUserUtil.ValidUser(userID)

		balance, bErr := ctrl.service.GetBalance(userID)
		if uErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if bErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": bErr.Error()})
		} else if rType == "balance" {
			c.JSON(http.StatusOK, balance)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "query param error"})
		}
	})
}