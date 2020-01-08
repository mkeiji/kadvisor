package controllers

import (
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
)

type LoginController struct {
	loginService services.LoginService
}

func (l *LoginController) LoadEndpoints(router *gin.Engine) {
	// login(/login)
	router.POST("/api/login", func (context *gin.Context) {
		var enteredLogin structs.Login
		context.BindJSON(&enteredLogin)

		storedLogin := l.loginService.GetOneByEmail(enteredLogin.Email)

		if storedLogin.ID != 0 {
			isValidPassword := KeiPassUtil.IsValidPassword(storedLogin.Password, enteredLogin.Password)
			if isValidPassword {
				updatedLogin := l.loginService.UpdateLoginStatus(storedLogin, true)
				context.JSON(http.StatusOK, gin.H{"login": updatedLogin})
			} else {
				context.JSON(http.StatusBadRequest, gin.H{"error": "ERROR: wrong password"})
			}
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": "ERROR: email not found"})
		}
	})
}