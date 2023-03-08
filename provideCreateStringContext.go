package goCommsDefinitions

import "go.uber.org/fx"

func ProvideStringContext(name string, value string) fx.Option {
	return fx.Provide(
		fx.Annotated{
			Name: name,
			Target: func() string {
				return value
			},
		},
	)
}
