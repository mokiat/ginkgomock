package ginkgomock

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo"
)

type Context struct {
	ctx        context.Context
	controller *gomock.Controller
}

func (c *Context) Context() context.Context {
	return c.ctx
}

func (c *Context) Controller() *gomock.Controller {
	return c.controller
}

func Describe(text string, body func(*Context)) bool {
	mockCtx := &Context{}

	return ginkgo.Describe(text, func() {
		ginkgo.BeforeEach(func() {
			mockCtx.controller, mockCtx.ctx = gomock.WithContext(context.Background(), ginkgo.GinkgoT())
		})

		ginkgo.Describe("", func() {
			body(mockCtx)
		})

		ginkgo.AfterEach(func() {
			mockCtx.controller.Finish()
		})
	})
}
