package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
)

type LoginController struct {
	loginservice services.LoginService
}

func (l *LoginController) LoadEndpoints(router *gin.Engine) {
	// login(/login)
	router.POST("/api/login", func (context *gin.Context) {
		var enteredLogin structs.Login
		context.BindJSON(&enteredLogin)

		storedLogin := l.loginservice.GetOneByEmail(enteredLogin.Email)

		isValidPassword := KeiPassUtil.IsValidPassword(storedLogin.Password, enteredLogin.Password)

		if isValidPassword {
			context.JSON(http.StatusOK, gin.H{"login": "user CAN login"})
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": "ERROR: wrong password"})
		}
	})
}