package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiUserUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
	"strconv"
)

type SubClassController struct {
	service 	services.SubClassService
}

func (ctrl *SubClassController) LoadEndpoints(router *gin.Engine) {
	// get(/subclass?id?classid)
	router.GET("/api/kadvisor/:uid/subclass", func (context *gin.Context) {
		subClassID	, _ := strconv.Atoi(context.Query("id"))
		classID		, _ := strconv.Atoi(context.Query("classid"))
		userID		, _ := strconv.Atoi(context.Param("uid"))

		uErr := KeiUserUtil.ValidUser(userID)

		getSubclassById := subClassID != 0 && classID == 0
		getSubclassesByUid := subClassID == 0 && classID == 0
		getSubclassesByClassid := subClassID == 0 && classID != 0

		if uErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if getSubclassById {
			subclass, err := ctrl.service.GetOneById(subClassID)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"subClass": subclass})
			}
		} else if getSubclassesByUid {
			subclasses, err := ctrl.service.GetManyByUserId(userID)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"subClasses": subclasses})
			}
		} else if getSubclassesByClassid {
			subclasses, err := ctrl.service.GetManyByClassId(classID)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, gin.H{"subClasses": subclasses})
			}
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": "query param error"})
		}
	})

	// post(/subclass)
	router.POST("/api/kadvisor/:uid/subclass", func (context *gin.Context) {
		var subclass structs.SubClass
		context.BindJSON(&subclass)

		userID, _ := strconv.Atoi(context.Param("uid"))
		uErr := KeiUserUtil.ValidUser(userID)

		saved, err := ctrl.service.Post(subclass)
		if uErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"subclass": saved})
		}
	})

	// delete(/subclass?id)
	router.DELETE("/api/kadvisor/:uid/subclass", func (context *gin.Context) {
		userID		, _ := strconv.Atoi(context.Param("uid"))
		subClassID	, _ := strconv.Atoi(context.Query("id"))

		uErr := KeiUserUtil.ValidUser(userID)

		deletedID, err := ctrl.service.Delete(subClassID)
		if uErr != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": uErr.Error()})
		} else if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusOK, gin.H{"deletedID": deletedID})
		}
	})
}