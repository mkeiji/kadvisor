package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
	"kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EntryController struct {
	Service           i.EntryService
	UsrService        i.UserService
	Auth              i.KeiAuthService
	ValidationService i.ValidationService
}

func NewEntryController() EntryController {
	return EntryController{
		Service:           s.NewEntryService(),
		UsrService:        s.NewUserService(),
		Auth:              s.NewKeiAuthService(),
		ValidationService: s.NewValidationService(),
	}
}

func (this EntryController) LoadEndpoints(router *gin.Engine) {
	this.Service = s.NewEntryService()
	this.UsrService = s.NewUserService()

	entryRoutes := router.Group("/api/kadvisor/:uid")
	permission := enums.REGULAR
	jwt, err := this.Auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	entryRoutes.Use(jwt.MiddlewareFunc())
	{
		entryRoutes.GET("/entry", this.GetEntry)
		entryRoutes.POST("/entry", this.PostEntry)
		entryRoutes.PUT("/entry", this.PutEntry)
		entryRoutes.DELETE("/entry", this.DeleteEntry)
	}
}

// get(/entry?id?classid?limit)
func (this EntryController) GetEntry(c *gin.Context) {
	var response dtos.KhttpResponse
	userID, _ := strconv.Atoi(c.Param("uid"))
	id, _ := strconv.Atoi(c.Query("id"))
	classID, _ := strconv.Atoi(c.Query("classid"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	getEntryById := id != 0 && classID == 0
	getEntriesByClassId := id == 0 && classID != 0 && limit != 0
	getEntryByUserId := id == 0 && classID == 0 && limit != 0

	if getEntryByUserId {
		response = this.Service.GetManyByUserId(userID, limit)
	} else if getEntryById {
		response = this.Service.GetOneById(id)
	} else if getEntriesByClassId {
		response = this.Service.GetManyByClassId(classID, limit)
	}

	c.JSON(response.Status, response.Body)
	return
}

// post(/entry)
func (this EntryController) PostEntry(c *gin.Context) {
	var response dtos.KhttpResponse
	var entry structs.Entry

	userID, _ := strconv.Atoi(c.Param("uid"))
	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	c.ShouldBindJSON(&entry)
	response = this.ValidationService.GetResponse(
		validators.NewEntryValidator(),
		entry,
	)
	if u.IsOKresponse(response.Status) {
		response = this.Service.Post(entry)
	}

	c.JSON(response.Status, response.Body)
	return
}

// put(/entry)
func (this EntryController) PutEntry(c *gin.Context) {
	var response dtos.KhttpResponse
	var entry structs.Entry

	userID, _ := strconv.Atoi(c.Param("uid"))
	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	c.ShouldBindJSON(&entry)
	response = this.ValidationService.GetResponse(
		validators.NewEntryValidator(),
		entry,
	)
	if u.IsOKresponse(response.Status) {
		response = this.Service.Put(entry)
	}

	c.JSON(response.Status, response.Body)
	return
}

// delete(/entry?id)
func (this EntryController) DeleteEntry(c *gin.Context) {
	var response dtos.KhttpResponse

	entryID, _ := strconv.Atoi(c.Query("id"))
	userID, _ := strconv.Atoi(c.Param("uid"))

	response = this.UsrService.GetOne(userID, false)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	}

	response = this.Service.Delete(entryID)
	c.JSON(response.Status, response.Body)
	return
}
