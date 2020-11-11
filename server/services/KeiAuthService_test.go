package services_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/resources/enums"
	"kadvisor/server/services"

	jwt "github.com/appleboy/gin-jwt/v2"
)

var _ = Describe("KeiAuthService", func() {
	Describe("GetAuthUtil", func() {
		It("should return an instance of jwt.GinJWTMiddleware", func() {
			service := services.KeiAuthService{}
			permission := enums.REGULAR
			expected := jwt.GinJWTMiddleware{
				Realm:             "test zone",
				SigningAlgorithm:  "HS256",
				Key:               []byte("secret key"),
				Timeout:           3600000000000,
				MaxRefresh:        3600000000000,
				IdentityKey:       "email",
				TokenLookup:       "header: Authorization, query: token, cookie: jwt",
				TokenHeadName:     "Bearer",
				SendCookie:        false,
				SecureCookie:      false,
				CookieHTTPOnly:    false,
				SendAuthorization: false,
				DisabledAbort:     false,
				CookieName:        "jwt",
				CookieSameSite:    0,
			}

			result, _ := service.GetAuthUtil(permission)
			Expect(result.Realm).To(Equal(expected.Realm))
			Expect(result.SigningAlgorithm).To(Equal(expected.SigningAlgorithm))
			Expect(result.Key).To(Equal(expected.Key))
			Expect(result.Timeout).To(Equal(expected.Timeout))
			Expect(result.MaxRefresh).To(Equal(expected.MaxRefresh))
			Expect(result.IdentityKey).To(Equal(expected.IdentityKey))
			Expect(result.TokenLookup).To(Equal(expected.TokenLookup))
			Expect(result.TokenHeadName).To(Equal(expected.TokenHeadName))
			Expect(result.SendCookie).To(Equal(expected.SendCookie))
			Expect(result.SecureCookie).To(Equal(expected.SecureCookie))
			Expect(result.CookieHTTPOnly).To(Equal(expected.CookieHTTPOnly))
			Expect(result.SendAuthorization).To(Equal(expected.SendAuthorization))
			Expect(result.DisabledAbort).To(Equal(expected.DisabledAbort))
			Expect(result.CookieName).To(Equal(expected.CookieName))
			Expect(result.CookieSameSite).To(Equal(expected.CookieSameSite))
		})
	})
})
