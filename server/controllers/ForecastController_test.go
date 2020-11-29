package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	g "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/controllers"
	"kadvisor/server/controllers/ControllerTestHelper"
	u "kadvisor/server/libs/KeiGenUtil"
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/interfaces/mocks"
	s "kadvisor/server/repository/structs"
	svc "kadvisor/server/services"
)

var _ = Describe("ForecastController", func() {
	const (
		testUserID = 1
	)

	var (
		controller        controllers.ForecastController
		mockCtrl          *g.Controller
		mockFcService     *mocks.MockForecastService
		mockUsrSvc        *mocks.MockUserService
		mockAuthSvc       *mocks.MockKeiAuthService
		mockValidationSvc *mocks.MockValidationService
		w                 *httptest.ResponseRecorder
		c                 *gin.Context
		r                 *gin.Engine
		endpoint          string
		route             string
	)

	BeforeEach(func() {
		endpoint = "/forecastTest"
		route = ControllerTestHelper.GetRoute(endpoint)
		w, c, r = ControllerTestHelper.SetupGinTest()
		mockCtrl = g.NewController(GinkgoT())
		mockFcService = mocks.NewMockForecastService(mockCtrl)
		mockAuthSvc = mocks.NewMockKeiAuthService(mockCtrl)
		mockUsrSvc = mocks.NewMockUserService(mockCtrl)
		mockValidationSvc = mocks.NewMockValidationService(mockCtrl)
		controller = controllers.ForecastController{
			FcService:         mockFcService,
			Auth:              mockAuthSvc,
			UsrService:        mockUsrSvc,
			ValidationService: mockValidationSvc,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := controllers.ForecastController{
				FcService:         svc.NewForecastService(),
				UsrService:        svc.NewUserService(),
				Auth:              svc.NewKeiAuthService(),
				ValidationService: svc.NewValidationService(),
			}

			Expect(controllers.NewForecastController()).To(Equal(expected))
		})
	})

	Describe("LoadEndpoints", func() {
		It("should set routes", func() {
			testJwt, _ := jwt.New(&jwt.GinJWTMiddleware{})
			expectedRoutes := []gin.RouteInfo{
				{
					Method: "GET",
					Path:   "/api/kadvisor/:uid/forecast",
				},
				{
					Method: "POST",
					Path:   "/api/kadvisor/:uid/forecast",
				},
				{
					Method: "DELETE",
					Path:   "/api/kadvisor/:uid/forecast",
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

	Describe("GetEntry", func() {
		var (
			request     *http.Request
			url         string
			queryParams map[string]string
		)

		Context("GetOneForecast", func() {
			It("should call fcService.GetOne and return an entries without error", func() {
				testID := 1
				thisYearInt := time.Now().Year()
				thisYearStr := u.ToString(thisYearInt)
				isPreloaded := false
				queryParams = map[string]string{"year": thisYearStr}
				url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				request, _ = http.NewRequest("GET", url, nil)
				testForecast := s.Forecast{Base: s.Base{ID: testID}}
				expectedBody := gin.H{
					"id":        float64(testID),
					"createdAt": "0001-01-01T00:00:00Z",
					"updatedAt": "0001-01-01T00:00:00Z",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockFcService.EXPECT().
					GetOne(testUserID, thisYearInt, isPreloaded).
					Return(dtos.NewKresponse(http.StatusOK, testForecast)).
					Times(1)

				r.GET(route, controller.GetOneForecast)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(result).To(Equal(expectedBody))
			})

			It("Invalid User - should return error if UserService.GetOne fails", func() {
				queryParams = map[string]string{"id": "1"}
				url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				request, _ = http.NewRequest("GET", url, nil)
				expected := gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusNotFound, expected)).
					Times(1)
				mockFcService.EXPECT().
					GetOne(g.Any(), g.Any(), g.Any()).
					Times(0)

				r.GET(route, controller.GetOneForecast)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(result).To(Equal(expected))
			})
		})

		Context("PostForecast", func() {
			var (
				reqBody      []byte
				expectedCode int
				expectedBody gin.H
			)

			BeforeEach(func() {
				reqBody = []byte("{}")
				url = ControllerTestHelper.GetUrl(testUserID, route)
			})

			AfterEach(func() {
				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expectedBody))
			})

			It("should call fcService.GetOne and return an entries without error", func() {
				testID := float64(1)
				testForecast := s.Forecast{}

				expectedCode = http.StatusOK
				request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
				expectedBody = gin.H{
					"id":        testID,
					"createdAt": "0001-01-01T00:00:00Z",
					"updatedAt": "0001-01-01T00:00:00Z",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(expectedCode, testForecast)).
					Times(1)
				mockFcService.EXPECT().
					Post(testForecast).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)

				r.POST(route, controller.PostForecast)
				r.ServeHTTP(c.Writer, request)
			})

			It("Invalid User - should return error if UserService.Post fails", func() {
				expectedCode = http.StatusNotFound

				request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
				expectedBody = gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Times(0)
				mockFcService.EXPECT().
					Post(g.Any()).
					Times(0)

				r.POST(route, controller.PostForecast)
				r.ServeHTTP(c.Writer, request)
			})

			It("Validation error - should return error if ValidationService fails", func() {
				expectedCode = http.StatusBadRequest

				request, _ = http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
				expectedBody = gin.H{
					"error": "test validation error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusOK, g.Any())).
					Times(1)
				mockValidationSvc.EXPECT().
					GetResponse(g.Any(), g.Any()).
					Return(dtos.NewKresponse(expectedCode, expectedBody)).
					Times(1)
				mockFcService.EXPECT().
					Post(g.Any()).
					Times(0)

				r.POST(route, controller.PostForecast)
				r.ServeHTTP(c.Writer, request)
			})
		})

		Context("DeleteForecast", func() {
			It("should call FcService.Delete and return deleted forecast", func() {
				expectedCode := http.StatusOK
				testDeletedID := 1
				expected := gin.H{"id": float64(testDeletedID)}
				queryParams = map[string]string{"id": "1"}
				url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				request, _ = http.NewRequest("DELETE", url, nil)

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(expectedCode, g.Any())).
					Times(1)
				mockFcService.EXPECT().
					Delete(testDeletedID).
					Return(dtos.NewKresponse(expectedCode, expected)).
					Times(1)

				r.DELETE(route, controller.DeleteForecast)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(expectedCode))
				Expect(result).To(Equal(expected))
			})

			It("Invalid User - should return error if UserService.Delete fails", func() {
				queryParams = map[string]string{"id": "1"}
				url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
				request, _ = http.NewRequest("DELETE", url, nil)
				expected := gin.H{
					"error": "test user error",
				}

				mockUsrSvc.EXPECT().
					GetOne(testUserID, false).
					Return(dtos.NewKresponse(http.StatusNotFound, expected)).
					Times(1)
				mockFcService.EXPECT().
					Delete(g.Any()).
					Times(0)

				r.DELETE(route, controller.DeleteForecast)
				r.ServeHTTP(c.Writer, request)

				var result gin.H
				json.Unmarshal(w.Body.Bytes(), &result)

				Expect(w.Code).To(Equal(http.StatusNotFound))
				Expect(result).To(Equal(expected))
			})
		})
	})
})
