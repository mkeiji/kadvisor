package controllers

import (
	"github.com/gin-gonic/gin"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"
)

type LookupController struct {
	Service    s.LookupService
	UsrService s.UserService
	Auth       s.KeiAuthService
}

func (ctrl LookupController) LoadEndpoints(router *gin.Engine) {
	ctrl.Service = s.NewLookupService()
	ctrl.UsrService = s.NewUserService()

	lookupRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	lookupRoutes.Use(jwt.MiddlewareFunc())
	{
		// get(/lookup?codeGroup)
		lookupRoutes.GET("/lookup", func(c *gin.Context) {
			var response dtos.KhttpResponse

			userID, _ := strconv.Atoi(c.Param("uid"))
			codeGroup := c.Query("codeGroup")

			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			response = ctrl.Service.GetAllByCodeGroup(codeGroup)
			c.JSON(response.Status, response.Body)
			return
		})
	}
}
