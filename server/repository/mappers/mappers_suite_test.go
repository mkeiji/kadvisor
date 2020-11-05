package mappers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMappers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mappers Suite")
}
