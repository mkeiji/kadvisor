package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassController struct {
	Service           s.ClassService
	Auth              s.KeiAuthService
	UsrService        s.UserService
	ValidationService s.ValidationService
}

func (ctrl ClassController) LoadEndpoints(router *gin.Engine) {
	ctrl.Service = s.NewClassService()
	ctrl.UsrService = s.NewUserService()

	classRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	classRoutes.Use(jwt.MiddlewareFunc())
	{
		// getOne(/class?id)
		classRoutes.GET("/class", func(c *gin.Context) {
			var response dtos.KhttpResponse
			userID, _ := strconv.Atoi(c.Param("uid"))
			classID, _ := strconv.Atoi(c.Query("id"))

			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			response = ctrl.Service.GetClass(userID, classID)
			c.JSON(response.Status, response.Body)
			return
		})

		// post(/class)
		classRoutes.POST("/class", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var class structs.Class

			userID, _ := strconv.Atoi(c.Param("uid"))
			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&class)
			response = ctrl.ValidationService.GetResponse(
				v.NewClassValidator(),
				class,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.Service.Post(class)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// put(/class)
		classRoutes.PUT("/class", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var class structs.Class

			userID, _ := strconv.Atoi(c.Param("uid"))
			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			c.BindJSON(&class)
			response = ctrl.ValidationService.GetResponse(
				v.NewClassValidator(),
				class,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.Service.Put(class)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// delete(/class?id)
		classRoutes.DELETE("/class", func(c *gin.Context) {
			var response dtos.KhttpResponse

			classID, _ := strconv.Atoi(c.Query("id"))
			userID, _ := strconv.Atoi(c.Param("uid"))

			response = ctrl.UsrService.GetOne(userID, false)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			}

			response = ctrl.Service.Delete(classID)
			c.JSON(response.Status, response.Body)
			return
		})
	}
}
