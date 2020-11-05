package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	"kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntryController struct {
	service           services.EntryService
	usrService        services.UserService
	auth              services.KeiAuthService
	validationService services.ValidationService
}

func (ctrl *EntryController) LoadEndpoints(router *gin.Engine) {
	entryRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	entryRoutes.Use(jwt.MiddlewareFunc())
	{
		// get(/entry?id?classid?limit)
		entryRoutes.GET("/entry", func(c *gin.Context) {
			var response dtos.KhttpResponse
			userID, _ := strconv.Atoi(c.Param("uid"))
			id, _ := strconv.Atoi(c.Query("id"))
			classID, _ := strconv.Atoi(c.Query("classid"))
			limit, _ := strconv.Atoi(c.Query("limit"))

			response = ctrl.usrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			getEntryById := id != 0 && classID == 0
			getEntriesByClassId := id == 0 && classID != 0
			getEntryByUserId := id == 0 && classID == 0

			if getEntryByUserId {
				response = ctrl.service.GetManyByUserId(userID, limit)
			} else if getEntryById {
				response = ctrl.service.GetOneById(id)
			} else if getEntriesByClassId {
				response = ctrl.service.GetManyByClassId(classID, limit)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// post(/entry)
		entryRoutes.POST("/entry", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var entry structs.Entry

			userID, _ := strconv.Atoi(c.Param("uid"))
			response = ctrl.usrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&entry)
			response = ctrl.validationService.GetResponse(
				validators.NewEntryValidator(),
				entry,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.service.Post(entry)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// put(/entry)
		entryRoutes.PUT("/entry", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var entry structs.Entry

			userID, _ := strconv.Atoi(c.Param("uid"))
			response = ctrl.usrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&entry)
			response = ctrl.validationService.GetResponse(
				validators.NewEntryValidator(),
				entry,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.service.Put(entry)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// delete(/entry?id)
		entryRoutes.DELETE("/entry", func(c *gin.Context) {
			var response dtos.KhttpResponse

			entryID, _ := strconv.Atoi(c.Query("id"))
			userID, _ := strconv.Atoi(c.Param("uid"))

			response = ctrl.usrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			response = ctrl.service.Delete(entryID)
			c.JSON(response.Status, response.Body)
			return
		})
	}
}
