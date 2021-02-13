package apiTests_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authentication", func() {
	Describe("jwt", func() {
		const endpoint = "/api/refresh_token"

		It("no token - should return unauthorized response", func() {
			resp, respErr := http.Get(testServer.URL + endpoint)
			Expect(respErr).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).Should(Equal(http.StatusUnauthorized))
			defer resp.Body.Close()
		})

		It("refresh token - should return ok response", func() {
			req := kRequest("GET", endpoint, nil)
			resp := kSendAndAssert(req, http.StatusOK)
			defer resp.Body.Close()
		})
	})
})
