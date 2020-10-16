package controllers

import (
	"errors"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginController class
type LoginController struct {
	loginService services.LoginService
	auth         services.KeiAuthService
}

// LoadEndpoints enpoints list
func (ctrl *LoginController) LoadEndpoints(router *gin.Engine) {
	loginRoutes := router.Group("/api")
	permission := enums.REGULAR
	jwt, err := ctrl.auth.GetAuthUtil(permission)
	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	// authentication(/auth)
	router.POST("/api/auth", jwt.LoginHandler)

	// refreshToken(/refresh_token)
	router.GET("/api/refresh_token", jwt.RefreshHandler)

	// protected
	loginRoutes.Use(jwt.MiddlewareFunc())
	{
		// login(/login)
		loginRoutes.POST("/login", func(context *gin.Context) {
			var enteredLogin structs.Login
			context.BindJSON(&enteredLogin)

			storedLogin, err := ctrl.loginService.GetOneByEmail(enteredLogin.Email)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				if KeiPassUtil.IsValidPassword(storedLogin.Password, enteredLogin.Password) {
					updatedLogin, err := ctrl.loginService.UpdateLoginStatus(storedLogin, true)
					if err != nil {
						context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					} else {
						context.JSON(http.StatusOK, updatedLogin)
					}
				} else {
					context.JSON(
						http.StatusBadRequest,
						gin.H{"error": errors.New("wrong password").Error()},
					)
				}
			}
		})

		// login(/login)
		loginRoutes.PUT("/login", func(context *gin.Context) {
			var login structs.Login
			context.BindJSON(&login)

			updated, err := ctrl.loginService.Put(login)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				context.JSON(http.StatusOK, updated)
			}
		})

		//logout(/logout)
		loginRoutes.POST("/logout", func(context *gin.Context) {
			var currentLogin structs.Login
			context.BindJSON(&currentLogin)

			storedLogin, err := ctrl.loginService.GetOneByEmail(currentLogin.Email)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				if storedLogin.IsLoggedIn == true {
					updatedLogin, err := ctrl.loginService.UpdateLoginStatus(storedLogin, false)
					if err != nil {
						context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					} else {
						context.JSON(http.StatusOK, updatedLogin)
					}
				}
			}
		})
	}
}
