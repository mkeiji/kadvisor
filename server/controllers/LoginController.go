package controllers

import (
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	i "kadvisor/server/repository/interfaces"
	"kadvisor/server/repository/structs"
	v "kadvisor/server/repository/validators"
	"kadvisor/server/resources/enums"
	s "kadvisor/server/services"
	"log"

	"github.com/gin-gonic/gin"
)

// LoginController class
type LoginController struct {
	LoginService      i.LoginService
	Auth              i.KeiAuthService
	ValidationService i.ValidationService
}

func NewLoginController() LoginController {

	return LoginController{
		LoginService:      s.NewLoginService(),
		Auth:              s.NewKeiAuthService(),
		ValidationService: s.NewValidationService(),
	}
}

// LoadEndpoints enpoints list
func (this LoginController) LoadEndpoints(router *gin.Engine) {
	this.LoginService = s.NewLoginService()

	loginRoutes := router.Group("/api")
	minPermission := enums.REGULAR
	jwt, err := this.Auth.GetAuthUtil(minPermission)
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
		loginRoutes.POST("/login", this.PostLogin)
		loginRoutes.PUT("/login", this.PutLogin)
		loginRoutes.POST("/logout", this.PostLogout)
	}
}

// login(/login)
func (this LoginController) PostLogin(c *gin.Context) {
	var response dtos.KhttpResponse
	var enteredLogin structs.Login

	c.ShouldBindJSON(&enteredLogin)
	response = this.ValidationService.GetResponse(
		v.NewLoginValidator(),
		enteredLogin,
	)
	if u.IsOKresponse(response.Status) {
		response = this.LoginService.UpdateLoginStatus(enteredLogin, true)
	}

	c.JSON(response.Status, response.Body)
	return
}

// login(/login)
func (this LoginController) PutLogin(c *gin.Context) {
	var response dtos.KhttpResponse
	var login structs.Login

	c.ShouldBindJSON(&login)
	response = this.LoginService.Put(login)
	c.JSON(response.Status, response.Body)
	return
}

//logout(/logout)
func (this LoginController) PostLogout(c *gin.Context) {
	var response dtos.KhttpResponse
	var currentLogin structs.Login
	c.ShouldBindJSON(&currentLogin)

	response = this.LoginService.GetOneByEmail(currentLogin.Email)
	if !u.IsOKresponse(response.Status) {
		c.JSON(response.Status, response.Body)
		return
	} else {
		storedLogin := response.Body.(structs.Login)
		if storedLogin.IsLoggedIn == true {
			response = this.LoginService.UpdateLoginStatus(storedLogin, false)
		}
	}

	c.JSON(response.Status, response.Body)
	return
}
