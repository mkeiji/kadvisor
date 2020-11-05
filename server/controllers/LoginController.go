package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	"kadvisor/server/services"
	"log"

	"github.com/gin-gonic/gin"
)

// LoginController class
type LoginController struct {
	loginService      services.LoginService
	auth              services.KeiAuthService
	validationService services.ValidationService
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
		loginRoutes.POST("/login", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var enteredLogin structs.Login

			c.BindJSON(&enteredLogin)
			response = ctrl.validationService.GetResponse(
				v.NewLoginValidator(),
				enteredLogin,
			)
			if u.IsOKresponse(response.Status) {
				response = ctrl.loginService.UpdateLoginStatus(enteredLogin, true)
			}

			c.JSON(response.Status, response.Body)
			return
		})

		// login(/login)
		loginRoutes.PUT("/login", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var login structs.Login

			c.BindJSON(&login)
			response = ctrl.loginService.Put(login)
			c.JSON(response.Status, response.Body)
			return
		})

		//logout(/logout)
		loginRoutes.POST("/logout", func(c *gin.Context) {
			var response dtos.KhttpResponse
			var currentLogin structs.Login
			c.BindJSON(&currentLogin)

			response = ctrl.loginService.GetOneByEmail(currentLogin.Email)
			if !u.IsOKresponse(response.Status) {
				c.JSON(response.Status, response.Body)
				return
			} else {
				storedLogin := response.Body.(structs.Login)
				if storedLogin.IsLoggedIn == true {
					response = ctrl.loginService.UpdateLoginStatus(storedLogin, false)
				}
			}

			c.JSON(response.Status, response.Body)
			return
		})
	}
}
