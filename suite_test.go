package ginkgomock_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgomock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgomock Suite")
}
