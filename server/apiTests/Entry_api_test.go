package apiTests_test

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	s "kadvisor/server/repository/structs"
)

var _ = Describe("EntryApi", func() {
	const (
		ENTRY_ENDPOINT   = "/api/kadvisor/:uid/entry"
		CLASS_ENDPOINT   = "/api/kadvisor/:uid/class"
		TEST_CLASS_NAME  = "testClass"
		TEST_DESCRIPTION = "testDescription"
		TEST_OBS         = "testObs"
	)
	var (
		testClass s.Class
		testEntry s.Entry
		today     time.Time
	)

	buildTestClass := func() s.Class {
		return s.Class{
			UserID:      testUserRegular.Login.UserID,
			Name:        TEST_CLASS_NAME,
			Description: TEST_DESCRIPTION,
		}
	}

	buildTestEntry := func() s.Entry {
		return s.Entry{
			UserID:          testUserRegular.Login.UserID,
			ClassID:         testClass.Base.ID,
			EntryTypeCodeID: "INCOME_ENTRY_TYPE",
			Date:            today,
			Amount:          float64(10),
			Description:     TEST_DESCRIPTION,
			Obs:             TEST_OBS,
		}
	}

	postTestClass := func() {
		respStatus, postClassErr := kMakeRequest("POST", CLASS_ENDPOINT, buildTestClass(), &testClass, nil)
		if postClassErr != nil {
			panic(fmt.Sprintf("\nError: %v", postClassErr))
		}

		Expect(respStatus).To(Equal(http.StatusOK))
	}

	postTestEntry := func() {
		respStatus, postEntryErr := kMakeRequest("POST", ENTRY_ENDPOINT, buildTestEntry(), &testEntry, nil)
		if postEntryErr != nil {
			panic(fmt.Sprintf("\nError: %v", postEntryErr))
		}

		Expect(respStatus).To(Equal(http.StatusOK))
	}

	BeforeEach(func() {
		today = time.Now()
		postTestClass()
	})

	Describe("GetEntry", func() {
		BeforeEach(func() {
			postTestEntry()
		})

		Context("No error", func() {
			It("should get entry with status ok", func() {
				params := map[string]string{"id": strconv.Itoa(testEntry.Base.ID)}

				var savedEntry s.Entry
				respStatus, getEntryErr := kMakeRequest(
					"GET", ENTRY_ENDPOINT, nil, &savedEntry, params,
				)

				Expect(getEntryErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(savedEntry).To(Equal(testEntry))
			})
		})

		Context("Error", func() {
			It("should return bad request for invalid id", func() {
				invalidID := 9999
				expectedErrMsg := "record not found"
				params := map[string]string{"id": strconv.Itoa(invalidID)}

				var savedEntry s.Entry
				respStatus, getEntryErr := kMakeRequest(
					"GET", ENTRY_ENDPOINT, nil, &savedEntry, params,
				)

				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(getEntryErr)).To(Equal(1))
				Expect(getEntryErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("PostEntry", func() {
		Context("No error", func() {
			It("should post an entry with ok response", func() {
				var entry s.Entry
				respStatus, postEntryErr := kMakeRequest(
					"POST", ENTRY_ENDPOINT, buildTestEntry(), &entry, nil,
				)
				Expect(postEntryErr).To(BeNil())
				Expect(respStatus).To(Equal(http.StatusOK))
			})
		})

		Context("Error", func() {
			It("should return bad request for invalid body", func() {
				invalidClassID := 9999
				invalidTestEntry := buildTestEntry()
				invalidTestEntry.ClassID = invalidClassID
				expectedErrMsg := "Key: 'Entry.ClassID' Error:Field validation for 'ClassID' invalid classID"

				var entry s.Entry
				respStatus, postEntryErr := kMakeRequest(
					"POST", ENTRY_ENDPOINT, invalidTestEntry, &entry, nil,
				)

				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(postEntryErr)).To(Equal(1))
				Expect(postEntryErr[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("PutForecastEntry", func() {
		const newAmount = float64(5)
		var (
			update s.Entry
		)
		BeforeEach(func() {
			postTestEntry()

			update = testEntry
			update.Amount = newAmount
		})

		Context("No error", func() {
			It("should return updated entry with ok response", func() {
				var result s.Entry
				respStatus, err := kMakeRequest("PUT", ENTRY_ENDPOINT, update, &result, nil)
				if err != nil {
					panic(fmt.Sprintf("\nError: %v", err))
				}

				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result.Base.ID).To(Equal(testEntry.Base.ID))
				Expect(result.Amount).To(Equal(newAmount))
			})
		})

		Context("Error", func() {
			It("should return bad request for invalid body", func() {
				invalidClassID := 9999
				expectedErrMsg := "Key: 'Entry.ClassID' Error:Field validation for 'ClassID' invalid classID"

				update.ClassID = invalidClassID

				var result s.Entry
				respStatus, err := kMakeRequest("PUT", ENTRY_ENDPOINT, update, &result, nil)

				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(err)).To(Equal(1))
				Expect(err[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})

	Describe("DeleteEntry", func() {
		BeforeEach(func() {
			postTestEntry()
		})

		Context("No error", func() {
			It("should return deleted entry ID with ok response", func() {
				params := map[string]string{"id": strconv.Itoa(testEntry.Base.ID)}

				var result int
				respStatus, err := kMakeRequest("DELETE", ENTRY_ENDPOINT, nil, &result, params)
				if err != nil {
					panic(fmt.Sprintf("\nError: %v", err))
				}

				Expect(respStatus).To(Equal(http.StatusOK))
				Expect(result).To(Equal(testEntry.Base.ID))
			})
		})

		Context("Error", func() {
			It("should return bad request if id is not valid", func() {
				invalidID := 9999
				expectedErrMsg := "record not found"
				params := map[string]string{"id": strconv.Itoa(invalidID)}

				var result int
				respStatus, err := kMakeRequest("DELETE", ENTRY_ENDPOINT, nil, &result, params)

				Expect(respStatus).To(Equal(http.StatusBadRequest))
				Expect(len(err)).To(Equal(1))
				Expect(err[0].Error()).To(Equal(expectedErrMsg))
			})
		})
	})
})
