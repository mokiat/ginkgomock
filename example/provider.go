package example

//go:generate mockgen -source provider.go -destination mock/ginkgomock_mocks.go -package mock

type Provider interface {
	Provide() string
}
