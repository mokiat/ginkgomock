# Ginkgo Mock

This package makes it easier to use [mock](https://github.com/golang/mock) with [ginkgo](https://github.com/onsi/ginkgo).

```go
var _ = ginkgomock.Describe("ControllerTest", func(c *ginkgomock.Context) { // note the context
  var controller *mocks.MockController

  BeforeEach(func() {
    controller = mocks.NewMockController(c.Controller())
  })

  // ...
})
```

There is no need to manage the lifecycle of the `gomock.Controller`. This is done automatically for you behind the scenes.
The `ginkgomock` library wraps your `Describe` with an outer `Describe` which has `BeforeEach` and `AfterEach` blocks
where the lifecycle of the mock controller is managed.
