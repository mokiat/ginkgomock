package ginkgomock_test

import (
	"github.com/mokiat/ginkgomock"
	"github.com/mokiat/ginkgomock/example/mock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = ginkgomock.Describe("GinkgoMock", func(c *ginkgomock.Context) {
	const testMessage = "mocking works"

	var provider *mock.MockProvider

	BeforeEach(func() {
		provider = mock.NewMockProvider(c.Controller())
		provider.EXPECT().Provide().Return(testMessage)
	})

	It("should succeed", func() {
		Expect(provider.Provide()).To(Equal(testMessage))
	})
})
