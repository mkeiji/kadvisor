package ControllerTestHelper

import (
	"fmt"
	. "github.com/onsi/gomega"
	"net/http/httptest"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const routerGroup = "/test/:uid"
const noUidRouterGroup = "/test"

func SetupGinTest() (*httptest.ResponseRecorder, *gin.Context, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	writer := httptest.NewRecorder()
	context, router := gin.CreateTestContext(writer)
	return writer, context, router
}

func GetRouteWithoutUid(endpoint string) string {
	return fmt.Sprintf("%v%v", noUidRouterGroup, endpoint)
}

func GetRoute(endpoint string) string {
	return fmt.Sprintf("%v%v", routerGroup, endpoint)
}

func GetUrl(userID int, route string) string {
	return GetUrlWithParams(userID, route, map[string]string{})
}

func GetUrlWithParams(userID int, route string, params map[string]string) string {
	var url string
	strID := strconv.Itoa(userID)

	if len(params) != 0 {
		url = fmt.Sprintf("%v?", route)

		i := 0
		for k, v := range params {
			var ampersend string
			if i == 0 {
				ampersend = ""
			} else {
				ampersend = "&"
			}

			url = fmt.Sprintf("%s%s%s=%s", url, ampersend, k, v)
			i++
		}
	} else {
		url = fmt.Sprintf("%v", route)
	}

	url = strings.Replace(url, ":uid", strID, -1)
	return url
}

func VerifyRoutes(resultRoutes []gin.RouteInfo, expectedRoutes []gin.RouteInfo) {
	Expect(len(resultRoutes)).To(Equal(len(expectedRoutes)))
	for i, route := range resultRoutes {
		Expect(route.Method).To(Equal(expectedRoutes[i].Method))
		Expect(route.Path).To(Equal(expectedRoutes[i].Path))
	}
}
