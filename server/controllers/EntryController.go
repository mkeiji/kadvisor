package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	"kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntryController struct {
	Service           s.EntryService
	UsrService        s.UserService
	Auth              s.KeiAuthService
	ValidationService s.ValidationService
}

func (ctrl EntryController) LoadEndpoints(router *gin.Engine) {
	ctrl.Service = s.NewEntryService()
	ctrl.UsrService = s.NewUserService()

	entryRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.Auth.GetAuthUtil(permission)
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

			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			getEntryById := id != 0 && classID == 0
			getEntriesByClassId := id == 0 && classID != 0
			getEntryByUserId := id == 0 && classID == 0

			if getEntryByUserId {
				response = ctrl.Service.GetManyByUserId(userID, limit)
			} else if getEntryById {
				response = ctrl.Service.GetOneById(id)
			} else if getEntriesByClassId {
				response = ctrl.Service.GetManyByClassId(classID, limit)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// post(/entry)
		entryRoutes.POST("/entry", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var entry structs.Entry

			userID, _ := strconv.Atoi(c.Param("uid"))
			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&entry)
			response = ctrl.ValidationService.GetResponse(
				validators.NewEntryValidator(),
				entry,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.Service.Post(entry)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// put(/entry)
		entryRoutes.PUT("/entry", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var entry structs.Entry

			userID, _ := strconv.Atoi(c.Param("uid"))
			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&entry)
			response = ctrl.ValidationService.GetResponse(
				validators.NewEntryValidator(),
				entry,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.Service.Put(entry)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// delete(/entry?id)
		entryRoutes.DELETE("/entry", func(c *gin.Context) {
			var response dtos.KhttpResponse

			entryID, _ := strconv.Atoi(c.Query("id"))
			userID, _ := strconv.Atoi(c.Param("uid"))

			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			response = ctrl.Service.Delete(entryID)
			c.JSON(response.Status, response.Body)
			return
		})
	}
}
