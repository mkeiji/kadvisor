package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
	"strconv"
)

type ForecastEntryController struct {
	service services.ForecastEntryService
}

func (ctrl *ForecastEntryController) LoadEndpoints(router *gin.Engine) {
	// put(/forecastentry)
	router.PUT("/api/kadvisor/:uid/forecastentry", func(c *gin.Context) {
		var entry structs.ForecastEntry

		userID, _ := strconv.Atoi(c.Param("uid"))
		uErr := KeiUserUtil.ValidUser(userID)

		c.BindJSON(&entry)
		updated, err := ctrl.service.Put(entry)
		if uErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, updated)
		}
	})
}
