package goCommsDefinitions

import "go.uber.org/fx"

func ProvideBool(name string, value bool) fx.Option {
	return fx.Provide(
		fx.Annotated{
			Name: name,
			Target: func() (bool, error) {
				return value, nil
			},
		},
	)
}
