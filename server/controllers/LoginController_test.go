package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/controllers"
	"kadvisor/server/controllers/ControllerTestHelper"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("LoginController", func() {
	const (
		testUserID = 1
	)

	var (
		controller        controllers.LoginController
		mockCtrl          *g.Controller
		mockLoginSvc      *mocks.MockLoginService
		mockAuthSvc       *mocks.MockKeiAuthService
		mockValidationSvc *mocks.MockValidationService
		w                 *httptest.ResponseRecorder
		c                 *gin.Context
		r                 *gin.Engine
		endpoint          string
		route             string
	)

	BeforeEach(func() {
		endpoint = "/loginTest"
		route = ControllerTestHelper.GetRoute(endpoint)
		w, c, r = ControllerTestHelper.SetupGinTest()
		mockCtrl = g.NewController(GinkgoT())
		mockLoginSvc = mocks.NewMockLoginService(mockCtrl)
		mockAuthSvc = mocks.NewMockKeiAuthService(mockCtrl)
		mockValidationSvc = mocks.NewMockValidationService(mockCtrl)
		controller = controllers.LoginController{
			LoginService:      mockLoginSvc,
			Auth:              mockAuthSvc,
			ValidationService: mockValidationSvc,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := controllers.LoginController{
				LoginService:      svc.NewLoginService(),
				Auth:              svc.NewKeiAuthService(),
				ValidationService: svc.NewValidationService(),
			}

			Expect(controllers.NewLoginController()).To(Equal(expected))
		})
	})

	Describe("LoadEndpoints", func() {
		It("should set routes", func() {
			testJwt, _ := jwt.New(&jwt.GinJWTMiddleware{})
			expectedRoutes := []gin.RouteInfo{
				{
					Method: "POST",
					Path:   "/api/login",
				},
				{
					Method: "POST",
					Path:   "/api/logout",
				},
				{
					Method: "POST",
					Path:   "/api/auth",
				},
				{
					Method: "GET",
					Path:   "/api/refresh_token",
				},
				{
					Method: "PUT",
					Path:   "/api/login",
				},
			}

			mockAuthSvc.EXPECT().
				GetAuthUtil(g.Any()).
				Return(testJwt, nil).
				Times(1)

			controller.LoadEndpoints(r)

			resultRoutes := r.Routes()
			ControllerTestHelper.VerifyRoutes(resultRoutes, expectedRoutes)
		})
	})

	Describe("Endpoints", func() {
		var (
			request      *http.Request
			url          string
			reqBody      []byte
			expectedCode int
			expectedBody gin.H
			testID       float64
		)

		BeforeEach(func() {
			testID = float64(1)

			reqBody = []byte("{}")
			url = ControllerTestHelper.GetUrl(testUserID, route)
		})

		AfterEach(func() {
			r.ServeHTTP(c.Writer, request)

			var result gin.H
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result).To(Equal(expectedBody))
		})

		It("PostLogin - should call service.Post and update Login status", func() {
			expectedCode = http.StatusOK
			request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
			expectedBody = gin.H{
				"id": testID,
			}

			mockValidationSvc.EXPECT().
				GetResponse(g.Any(), g.Any()).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockLoginSvc.EXPECT().
				UpdateLoginStatus(g.Any(), true).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)

			r.POST(route, controller.PostLogin)
		})

		It("PostLogin - should return error if validation service fails", func() {
			expectedCode = http.StatusBadRequest
			request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
			expectedBody = gin.H{
				"error": "test validation error",
			}

			mockValidationSvc.EXPECT().
				GetResponse(g.Any(), g.Any()).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockLoginSvc.EXPECT().
				UpdateLoginStatus(g.Any(), g.Any()).
				Times(0)

			r.POST(route, controller.PostLogin)
		})

		It("PutLogin - should call service.Put", func() {
			expectedCode = http.StatusOK
			request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
			expectedBody = gin.H{
				"id": testID,
			}
			mockLoginSvc.EXPECT().
				Put(g.Any()).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)

			r.PUT(route, controller.PutLogin)
		})

		It("PostLogout - should call service.UpdateLoginStatus", func() {
			testLogin := s.Login{
				Base:       s.Base{ID: 1},
				IsLoggedIn: true,
			}

			expectedCode = http.StatusOK
			request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
			expectedBody = gin.H{
				"id":         testID,
				"isLoggedIn": false,
			}

			mockLoginSvc.EXPECT().
				GetOneByEmail(g.Any()).
				Return(dtos.NewKresponse(expectedCode, testLogin)).
				Times(1)
			mockLoginSvc.EXPECT().
				UpdateLoginStatus(g.Any(), false).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)

			r.POST(route, controller.PostLogout)
		})

		It("PostLogout - should return error if getOneByEmail fails", func() {
			expectedCode = http.StatusNotFound
			request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
			expectedBody = gin.H{
				"error": "test service error",
			}

			mockLoginSvc.EXPECT().
				GetOneByEmail(g.Any()).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockLoginSvc.EXPECT().
				UpdateLoginStatus(g.Any(), g.Any()).
				Times(0)

			r.POST(route, controller.PostLogout)
		})
	})

	Describe("PostLogout", func() {
		It("does not update status if isLoggedIn already false", func() {
			testLogin := s.Login{
				IsLoggedIn: false,
			}
			reqBody := []byte("{}")
			url := ControllerTestHelper.GetUrl(testUserID, route)
			expectedCode := http.StatusOK
			request, _ := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))

			g.InOrder(
				mockLoginSvc.EXPECT().
					GetOneByEmail(g.Any()).
					Return(dtos.NewKresponse(expectedCode, testLogin)).
					Times(1),
				mockLoginSvc.EXPECT().
					UpdateLoginStatus(testLogin, false).
					Times(0),
			)

			r.POST(route, controller.PostLogout)
			r.ServeHTTP(c.Writer, request)

			var result s.Login
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result.IsLoggedIn).To(BeFalse())
		})
	})
})
