package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

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

var _ = Describe("UserController", func() {
	const (
		testUserID = 1
	)

	var (
		controller        controllers.UserController
		mockCtrl          *g.Controller
		mockUsrSvc        *mocks.MockUserService
		mockAuthSvc       *mocks.MockKeiAuthService
		mockValidationSvc *mocks.MockValidationService
		w                 *httptest.ResponseRecorder
		c                 *gin.Context
		r                 *gin.Engine
		endpoint          string
		route             string
		url               string
		queryParams       map[string]string
		request           *http.Request
		expectedCode      int
		expectedBody      gin.H
	)

	BeforeEach(func() {
		endpoint = "/userTest"
		route = ControllerTestHelper.GetRouteWithoutUid(endpoint)
		w, c, r = ControllerTestHelper.SetupGinTest()
		mockCtrl = g.NewController(GinkgoT())
		mockUsrSvc = mocks.NewMockUserService(mockCtrl)
		mockAuthSvc = mocks.NewMockKeiAuthService(mockCtrl)
		mockValidationSvc = mocks.NewMockValidationService(mockCtrl)
		controller = controllers.UserController{
			UserService:       mockUsrSvc,
			Auth:              mockAuthSvc,
			ValidationService: mockValidationSvc,
		}

		expectedCode = http.StatusOK
		expectedBody = gin.H{}

	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := controllers.UserController{
				UserService:       svc.NewUserService(),
				Auth:              svc.NewKeiAuthService(),
				ValidationService: svc.NewValidationService(),
			}

			Expect(controllers.NewUserController()).To(Equal(expected))
		})
	})

	Describe("LoadEndpoints", func() {
		It("should set routes", func() {
			testJwt, _ := jwt.New(&jwt.GinJWTMiddleware{})
			expectedRoutes := []gin.RouteInfo{
				{
					Method: "POST",
					Path:   "/api/user",
				},
				{
					Method: "GET",
					Path:   "/api/user/:id",
				},
				{
					Method: "GET",
					Path:   "/api/users",
				},
				{
					Method: "PUT",
					Path:   "/api/user",
				},
				{
					Method: "DELETE",
					Path:   "/api/user/:id",
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

	Describe("GET endpoints", func() {
		AfterEach(func() {
			r.ServeHTTP(c.Writer, request)

			var result gin.H
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result).To(Equal(expectedBody))
		})

		Context("GetOneUser", func() {
			It("No preload - should call UserService.GetOne with preloaded = false", func() {
				strTestUserID := fmt.Sprintf("%v", testUserID)
				route = fmt.Sprintf("%v%v", route, "/:id")
				url = ControllerTestHelper.GetUrl(testUserID, route)
				url = strings.Replace(url, ":id", strTestUserID, -1)
				request, _ = http.NewRequest("GET", url, nil)

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.GET(route, controller.GetOneUser)
			})

			It("Yes preload - should call UserService.GetOne with preloaded = true", func() {
				strTestUserID := fmt.Sprintf("%v", testUserID)
				route = fmt.Sprintf("%v%v", route, "/:id")
				queryParams = map[string]string{"preloaded": "true"}
				url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				url = strings.Replace(url, ":id", strTestUserID, -1)
				request, _ = http.NewRequest("GET", url, nil)

				mockUsrSvc.EXPECT().
					GetOne(testUserID, true).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.GET(route, controller.GetOneUser)
			})
		})

		Context("GetManyUsers", func() {
			It("No preload - should call UserService.GetOne with preloaded = false", func() {
				strTestUserID := fmt.Sprintf("%v", testUserID)
				route = fmt.Sprintf("%v%v", route, "/:id")
				url = ControllerTestHelper.GetUrl(testUserID, route)
				url = strings.Replace(url, ":id", strTestUserID, -1)
				request, _ = http.NewRequest("GET", url, nil)

				mockUsrSvc.EXPECT().
					GetMany(false).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.GET(route, controller.GetManyUsers)
			})

			It("Yes preload - should call UserService.GetOne with preloaded = true", func() {
				strTestUserID := fmt.Sprintf("%v", testUserID)
				route = fmt.Sprintf("%v%v", route, "/:id")
				queryParams = map[string]string{"preloaded": "true"}
				url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				url = strings.Replace(url, ":id", strTestUserID, -1)
				request, _ = http.NewRequest("GET", url, nil)

				mockUsrSvc.EXPECT().
					GetMany(true).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.GET(route, controller.GetManyUsers)
			})
		})
	})

	Describe("POST endpoints", func() {
		Context("PostUser", func() {
			var (
				reqBody []byte
				testID  float64
			)

			BeforeEach(func() {
				testID = float64(1)

				reqBody = []byte("{}")
				url = ControllerTestHelper.GetUrl(testUserID, route)
				request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
			})

			AfterEach(func() {
				r.POST(route, controller.PostUser)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})

			It("should call userService.Post", func() {
				expectedCode = http.StatusOK
				expectedBody = gin.H{
					"id": testID,
				}
				testUser := s.User{}

				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), testUser).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockUsrSvc.EXPECT().
					Post(g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
			})

			It("should return error if validation service fails", func() {
				expectedCode = http.StatusBadRequest
				expectedBody = gin.H{
					"error": "test validation error",
				}

				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
				mockUsrSvc.EXPECT().
					Post(g.Any()).
					Times(0)
			})
		})
	})

	Describe("PUT endpoints", func() {
		Context("UpdateUser", func() {
			var (
				reqBody []byte
				testID  float64
			)

			BeforeEach(func() {
				testID = float64(1)

				reqBody = []byte("{}")
				url = ControllerTestHelper.GetUrl(testUserID, route)
				request, _ = http.NewRequest("PUT", url, bytes.NewBuffer(reqBody))
			})

			AfterEach(func() {
				r.PUT(route, controller.UpdateUser)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})

			It("should call userService.Put", func() {
				expectedCode = http.StatusOK
				expectedBody = gin.H{
					"id": testID,
				}
				testUser := s.User{}

				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), testUser).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockUsrSvc.EXPECT().
					Put(testUser).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
			})

			It("should return error if validation service fails", func() {
				expectedCode = http.StatusBadRequest
				expectedBody = gin.H{
					"error": "test validation error",
				}

				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
				mockUsrSvc.EXPECT().
					Put(g.Any()).
					Times(0)
			})
		})
	})

	Describe("DELETE endpoint", func() {
		It("should call UserService.Delete", func() {
			strTestUserID := fmt.Sprintf("%v", testUserID)
			route = fmt.Sprintf("%v%v", route, "/:id")
			url = ControllerTestHelper.GetUrl(testUserID, route)
			url = strings.Replace(url, ":id", strTestUserID, -1)
			request, _ = http.NewRequest("DELETE", url, nil)

			mockUsrSvc.EXPECT().
				Delete(testUserID).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)

			r.DELETE(route, controller.DeleteUser)
			r.ServeHTTP(c.Writer, request)

			var result gin.H
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result).To(Equal(expectedBody))
		})
	})
})
