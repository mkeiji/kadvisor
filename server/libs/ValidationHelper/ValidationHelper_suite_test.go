package ValidationHelper_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidationHelper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ValidationHelper Suite")
}
