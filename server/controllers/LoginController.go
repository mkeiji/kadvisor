package controllers

import (
	"github.com/gin-gonic/gin"
	"errors"
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

		storedLogin, err := l.loginService.GetOneByEmail(enteredLogin.Email)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			if KeiPassUtil.IsValidPassword(storedLogin.Password, enteredLogin.Password) {
				updatedLogin, err := l.loginService.UpdateLoginStatus(storedLogin, true)
				if err != nil {
					context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
				} else {
					context.JSON(http.StatusOK, gin.H{"login": updatedLogin})
				}
			} else {
				context.JSON(http.StatusBadRequest, gin.H{"error": errors.New("wrong password").Error()})
			}
		}
	})
}