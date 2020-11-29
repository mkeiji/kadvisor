package controllers_test

import (
	"encoding/json"
	"fmt"
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
	"kadvisor/server/libs/dtos"
	"kadvisor/server/repository/interfaces/mocks"
	svc "kadvisor/server/services"
)

var _ = Describe("ReportController", func() {
	const (
		testUserID = 1
	)

	var (
		controller    controllers.ReportController
		mockCtrl      *g.Controller
		mockReportSvc *mocks.MockReportService
		mockUsrSvc    *mocks.MockUserService
		mockAuthSvc   *mocks.MockKeiAuthService
		w             *httptest.ResponseRecorder
		c             *gin.Context
		r             *gin.Engine
		endpoint      string
		route         string
		url           string
		queryParams   map[string]string
		request       *http.Request
		expectedCode  int
		expectedBody  gin.H
	)

	BeforeEach(func() {
		endpoint = "/reportTest"
		route = ControllerTestHelper.GetRoute(endpoint)
		w, c, r = ControllerTestHelper.SetupGinTest()
		mockCtrl = g.NewController(GinkgoT())
		mockReportSvc = mocks.NewMockReportService(mockCtrl)
		mockUsrSvc = mocks.NewMockUserService(mockCtrl)
		mockAuthSvc = mocks.NewMockKeiAuthService(mockCtrl)
		controller = controllers.ReportController{
			Service:    mockReportSvc,
			UsrService: mockUsrSvc,
			Auth:       mockAuthSvc,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("Constructor", func() {
		It("should return an instance", func() {
			expected := controllers.ReportController{
				Service:    svc.NewReportService(),
				UsrService: svc.NewUserService(),
				Auth:       svc.NewKeiAuthService(),
			}

			Expect(controllers.NewReportController()).To(Equal(expected))
		})
	})

	Describe("LoadEndpoints", func() {
		It("should set routes", func() {
			testJwt, _ := jwt.New(&jwt.GinJWTMiddleware{})
			expectedRoutes := []gin.RouteInfo{
				{
					Method: "GET",
					Path:   "/api/kadvisor/:uid/report",
				},
				{
					Method: "GET",
					Path:   "/api/kadvisor/:uid/reportavailable",
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

	Describe("GetReport", func() {
		var (
			typeBalance string
			typeYear    string
			typeYearFC  string
			thisYear    string
		)

		BeforeEach(func() {
			thisYear = fmt.Sprintf("%v", time.Now().Year())
			typeBalance = "BALANCE"
			typeYear = "YTD"
			typeYearFC = "YFC"
		})

		AfterEach(func() {
			r.GET(route, controller.GetReport)
			r.ServeHTTP(c.Writer, request)

			var result gin.H
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result).To(Equal(expectedBody))
		})

		It("should call service.GetBalance", func() {
			queryParams = map[string]string{"type": typeBalance}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)

			expectedCode = http.StatusOK
			expectedBody = gin.H{}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, g.Any())).
				Times(1)
			mockReportSvc.EXPECT().
				GetBalance(testUserID).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockReportSvc.EXPECT().
				GetYearToDateReport(g.Any(), g.Any()).
				Times(0)
			mockReportSvc.EXPECT().
				GetYearToDateWithForecastReport(g.Any(), g.Any()).
				Times(0)
		})

		It("should call service.GetYearToDateReport", func() {
			queryParams = map[string]string{"type": typeYear, "year": thisYear}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)

			expectedCode = http.StatusOK
			expectedBody = gin.H{}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, g.Any())).
				Times(1)
			mockReportSvc.EXPECT().
				GetBalance(g.Any()).
				Times(0)
			mockReportSvc.EXPECT().
				GetYearToDateReport(testUserID, g.Any()).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockReportSvc.EXPECT().
				GetYearToDateWithForecastReport(g.Any(), g.Any()).
				Times(0)
		})

		It("should call service.GetYearToDateWithForecastReport", func() {
			queryParams = map[string]string{"type": typeYearFC, "year": thisYear}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)

			expectedCode = http.StatusOK
			expectedBody = gin.H{}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, g.Any())).
				Times(1)
			mockReportSvc.EXPECT().
				GetBalance(g.Any()).
				Times(0)
			mockReportSvc.EXPECT().
				GetYearToDateReport(g.Any(), g.Any()).
				Times(0)
			mockReportSvc.EXPECT().
				GetYearToDateWithForecastReport(testUserID, g.Any()).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
		})

		It("should return error if UserService.GetOne fails", func() {
			queryParams = map[string]string{"type": typeBalance}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)

			expectedCode = http.StatusNotFound
			expectedBody = gin.H{
				"error": "test user error",
			}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockReportSvc.EXPECT().
				GetBalance(testUserID).
				Times(0)
			mockReportSvc.EXPECT().
				GetYearToDateReport(g.Any(), g.Any()).
				Times(0)
			mockReportSvc.EXPECT().
				GetYearToDateWithForecastReport(g.Any(), g.Any()).
				Times(0)
		})
	})

	Describe("GetReportAvailable", func() {
		AfterEach(func() {
			r.GET(route, controller.GetReportAvailable)
			r.ServeHTTP(c.Writer, request)

			var result gin.H
			json.Unmarshal(w.Body.Bytes(), &result)

			Expect(w.Code).To(Equal(expectedCode))
			Expect(result).To(Equal(expectedBody))
		})

		It("should call service.GetReportForecastAvailable", func() {
			queryParams = map[string]string{"forecast": "true"}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)

			expectedCode = http.StatusOK
			expectedBody = gin.H{}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, g.Any())).
				Times(1)
			mockReportSvc.EXPECT().
				GetReportForecastAvailable(testUserID).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
		})

		It("should call service.GetReportAvailable", func() {
			queryParams = map[string]string{"forecast": "false"}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)

			expectedCode = http.StatusOK
			expectedBody = gin.H{}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, g.Any())).
				Times(1)
			mockReportSvc.EXPECT().
				GetReportForecastAvailable(g.Any()).
				Times(0)
			mockReportSvc.EXPECT().
				GetReportAvailable(testUserID).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
		})

		It("should return error if UserService.GetOne fails", func() {
			queryParams = map[string]string{"forecast": "true"}
			url = ControllerTestHelper.GetUrlWithParams(testUserID, route, queryParams)
			request, _ = http.NewRequest("GET", url, nil)

			expectedCode = http.StatusNotFound
			expectedBody = gin.H{
				"error": "test user error",
			}

			mockUsrSvc.EXPECT().
				GetOne(testUserID, false).
				Return(dtos.NewKresponse(expectedCode, expectedBody)).
				Times(1)
			mockReportSvc.EXPECT().
				GetReportForecastAvailable(g.Any()).
				Times(0)
			mockReportSvc.EXPECT().
				GetReportAvailable(g.Any()).
				Times(0)
		})
	})
})
