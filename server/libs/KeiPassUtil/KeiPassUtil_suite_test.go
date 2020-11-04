package KeiPassUtil_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKeiPassUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KeiPassUtil Suite")
}
