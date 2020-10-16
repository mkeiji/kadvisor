package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"net/http"
	"strconv"
)

type EntryController struct {
	service services.EntryService
	auth    services.KeiAuthService
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
			userID, _ := strconv.Atoi(c.Param("uid"))
			id, _ := strconv.Atoi(c.Query("id"))
			classID, _ := strconv.Atoi(c.Query("classid"))
			limit, _ := strconv.Atoi(c.Query("limit"))

			uErr := KeiUserUtil.ValidUser(userID)

			getEntryById := id != 0 && classID == 0
			getEntriesByClassId := id == 0 && classID != 0
			getEntryByUserId := id == 0 && classID == 0

			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if getEntryByUserId {
				ctrl.getEntriesByUserId(c, userID, limit)
			} else if getEntryById {
				ctrl.getEntryById(c, id)
			} else if getEntriesByClassId {
				ctrl.getEntriesByClassId(c, classID, limit)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "query param error"})
			}
		})

		// post(/entry)
		entryRoutes.POST("/entry", func(c *gin.Context) {
			var entry structs.Entry

			userID, _ := strconv.Atoi(c.Param("uid"))
			uErr := KeiUserUtil.ValidUser(userID)

			c.BindJSON(&entry)
			saved, err := ctrl.service.Post(entry)
			if uErr != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
			} else if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, saved)
			}
		})

		// put(/entry)
		entryRoutes.PUT("/entry", func(c *gin.Context) {
			var entry structs.Entry

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

		// delete(/entry?id)
		entryRoutes.DELETE("/entry", func(c *gin.Context) {
			entryID, _ := strconv.Atoi(c.Query("id"))
			userID, _ := strconv.Atoi(c.Param("uid"))
			uErr := KeiUserUtil.ValidUser(userID)

			deletedID, err := ctrl.service.Delete(entryID)
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

func (ctrl *EntryController) getEntryById(c *gin.Context, entryID int) {
	entry, err := ctrl.service.GetOneById(entryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, entry)
	}
}

func (ctrl *EntryController) getEntriesByUserId(
	c *gin.Context, userID int, limit int) {
	entries, err := ctrl.service.GetManyByUserId(userID, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, entries)
	}
}

func (ctrl *EntryController) getEntriesByClassId(
	c *gin.Context, classID int, limit int) {
	entries, err := ctrl.service.GetManyByClassId(classID, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, entries)
	}
}
