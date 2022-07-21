package goCommsDefinitions

import (
	"go.uber.org/fx"
	"net/url"
)

func ProvideUrl(name string, connectionUrl *url.URL) fx.Option {
	return fx.Provide(
		fx.Annotated{
			Name: name,
			Target: func() *url.URL {
				return connectionUrl
			},
		},
	)
}
