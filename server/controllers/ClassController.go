package controllers

import (
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClassController struct {
	service services.ClassService
	auth    services.KeiAuthService
}

func (ctrl *ClassController) LoadEndpoints(router *gin.Engine) {
	classRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := ctrl.auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	classRoutes.Use(jwt.MiddlewareFunc())
	{
		// getOne(/class?id)
		classRoutes.GET("/class", func(c *gin.Context) {
			userID, _ := strconv.Atoi(c.Param("uid"))
			classID, _ := strconv.Atoi(c.Query("id"))

			uErr := KeiUserUtil.ValidUser(userID)

			getClassesByUserId := classID == 0 && userID != 0
			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if classID != 0 {
				class, err := ctrl.service.GetOneById(classID)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				} else {
					c.JSON(http.StatusOK, class)
				}
			} else if getClassesByUserId {
				classes, err := ctrl.service.GetManyByUserId(userID)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				} else {
					c.JSON(http.StatusOK, classes)
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "query param error"})
			}
		})

		// post(/class)
		classRoutes.POST("/class", func(c *gin.Context) {
			var class structs.Class

			userID, _ := strconv.Atoi(c.Param("uid"))
			uErr := KeiUserUtil.ValidUser(userID)

			c.BindJSON(&class)
			saved, err := ctrl.service.Post(class)

			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, saved)
			}
		})

		// put(/class)
		classRoutes.PUT("/class", func(c *gin.Context) {
			var class structs.Class

			userID, _ := strconv.Atoi(c.Param("uid"))
			uErr := KeiUserUtil.ValidUser(userID)

			c.BindJSON(&class)
			updated, err := ctrl.service.Put(class)
			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, updated)
			}
		})

		// delete(/class?id)
		classRoutes.DELETE("/class", func(c *gin.Context) {
			classID, _ := strconv.Atoi(c.Query("id"))
			userID, _ := strconv.Atoi(c.Param("uid"))
			uErr := KeiUserUtil.ValidUser(userID)

			deletedID, err := ctrl.service.Delete(classID)
			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, deletedID)
			}
		})
	}
}
