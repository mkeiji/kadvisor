package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
	"strconv"
)

type EntryController struct {
	service services.EntryService
}
func (ctrl *EntryController) LoadEndpoints(router *gin.Engine) {
	// get(/entry?id?classid?subclassid?)
	router.GET("/api/kadvisor/:uid/entry", func (c *gin.Context) {
		userID		, _ := strconv.Atoi(c.Param("uid"))
		id			, _ := strconv.Atoi(c.Query("id"))
		classID		, _ := strconv.Atoi(c.Query("classid"))
		subclassID	, _ := strconv.Atoi(c.Query("subclassid"))

		uErr := KeiUserUtil.ValidUser(userID)

		getEntryById 					:= id != 0 && classID == 0 && subclassID == 0
		getEntriesByUserId 				:= id == 0 && classID == 0 && subclassID == 0
		getEntriesByClassId 			:= id == 0 && classID != 0 && subclassID == 0
		getEntriesBySubclassId 			:= id == 0 && classID == 0 && subclassID != 0
		getEntriesByClassAndSubClass 	:= id == 0 && classID != 0 && subclassID != 0

		if uErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if getEntryById {
			_getEntryById(ctrl.service, c, id)
		} else if getEntriesByUserId {
			_getEntriesByUserId(ctrl.service, c, userID)
		} else if getEntriesByClassId {
			_getEntriesByClassId(ctrl.service, c, classID)
		} else if getEntriesBySubclassId {
			_getEntriesBySubclassId(ctrl.service, c, subclassID)
		} else if getEntriesByClassAndSubClass {
			_getEntriesByClassAndSubClass(ctrl.service, c, classID, subclassID)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "query param error"})
		}
	})

	// post(/entry)
	router.POST("/api/kadvisor/:uid/entry", func (c *gin.Context) {
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
	router.PUT("/api/kadvisor/:uid/entry", func (c *gin.Context) {
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
	router.DELETE("/api/kadvisor/:uid/entry", func (c *gin.Context) {
		entryID	, _ := strconv.Atoi(c.Query("id"))
		userID	, _ := strconv.Atoi(c.Param("uid"))
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

func _getEntryById(s services.EntryService, c *gin.Context, entryID int) {
	entry, err := s.GetOneById(entryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, entry)
	}
}

func _getEntriesByUserId(s services.EntryService, c *gin.Context, userID int) {
	entries, err := s.GetManyByUserId(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, entries)
	}
}

func _getEntriesByClassId(s services.EntryService, c *gin.Context, classID int) {
	entries, err := s.GetManyByClassId(classID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, entries)
	}
}

func _getEntriesBySubclassId(s services.EntryService, c *gin.Context, subclassID int) {
	entries, err := s.GetManyByClassId(subclassID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, entries)
	}
}

func _getEntriesByClassAndSubClass(
	s services.EntryService, c *gin.Context, classID int, subclassID int) {
	entries, err := s.GetManyByClassAndSubClassId(classID, subclassID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, entries)
	}
}