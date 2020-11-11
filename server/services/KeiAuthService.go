package services

import (
	"kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/KeiPassUtil"
	"kadvisor/server/repository/structs"
	"kadvisor/server/resources/enums"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var emailKey = "email"
var passwordKey = "password"

type login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type authLogin struct {
	Email    string
	Password string
}

type KeiAuthService struct {
	LoginService LoginService
}

func (svc KeiAuthService) GetAuthUtil(
	permissionLevel enums.RoleEnum,
) (*jwt.GinJWTMiddleware, error) {
	svc.LoginService = NewLoginService()

	var actualLogin structs.Login

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: emailKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*authLogin); ok {
				return jwt.MapClaims{
					emailKey:    v.Email,
					passwordKey: v.Password,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &authLogin{
				Email:    claims[emailKey].(string),
				Password: claims[passwordKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var claimLogin login
			if err := c.ShouldBind(&claimLogin); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			res := svc.LoginService.GetOneByEmail(claimLogin.Email)
			actualLogin := res.Body.(structs.Login)
			if svc.isValidLogin(actualLogin, claimLogin) {
				return &authLogin{
					Email:    claimLogin.Email,
					Password: claimLogin.Password,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			_, ok := data.(*authLogin)
			if ok && KeiGenUtil.HasPermission(
				actualLogin.RoleID, permissionLevel,
			) {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}

func (svc KeiAuthService) isValidLogin(actual structs.Login, claim login) bool {
	if actual.Email == claim.Email && KeiPassUtil.IsValidPassword(actual.Password, claim.Password) {
		return true
	} else {
		return false
	}
}
