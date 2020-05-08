package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/services"
	"net/http"
	"strconv"
)

type LookupController struct {
	service services.LookupService
}

func (ctrl * LookupController) LoadEndpoints(router *gin.Engine) {
	// get(/lookup?codeGroup)
	router.GET("/api/kadvisor/:uid/lookup", func (c *gin.Context) {
		userID, _ 	:= strconv.Atoi(c.Param("uid"))
		codeGroup	:= c.Query("codeGroup")

		uErr := KeiUserUtil.ValidUser(userID)

		lookups, err := ctrl.service.GetAllByCodeGroup(codeGroup)
		if uErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if codeGroup != "" {
			c.JSON(http.StatusOK, lookups)
		} else {
			c.JSON(http.StatusBadRequest,
				gin.H{"error": errors.New("missing codeGroup param")})
		}
	})
}