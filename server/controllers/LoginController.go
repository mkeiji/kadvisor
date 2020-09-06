package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/services"
	"net/http"
)

// LoginController class
type LoginController struct {
	loginService services.LoginService
}

// LoadEndpoints enpoints list
func (l *LoginController) LoadEndpoints(router *gin.Engine) {
	// login(/login)
	router.POST("/api/login",
		func(context *gin.Context) {

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
						context.JSON(http.StatusOK, updatedLogin)
					}
				} else {
					context.JSON(http.StatusBadRequest, gin.H{"error": errors.New("wrong password").Error()})
				}
			}
		})

	// login(/login)
	router.PUT("/api/login",
		func(context *gin.Context) {

			var login structs.Login
			context.BindJSON(&login)

			updated, err := l.loginService.Put(login)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, updated)
			}
		})

	//logout(/logout)
	router.POST("/api/logout", func(context *gin.Context) {
		var currentLogin structs.Login
		context.BindJSON(&currentLogin)

		storedLogin, err := l.loginService.GetOneByEmail(currentLogin.Email)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			if storedLogin.IsLoggedIn == true {
				updatedLogin, err := l.loginService.UpdateLoginStatus(storedLogin, false)
				if err != nil {
					context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				} else {
					context.JSON(http.StatusOK, updatedLogin)
				}
			}
		}
	})
}
