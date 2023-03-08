package goCommsDefinitions

import "go.uber.org/fx"

func ProvideTransportFactoryForWebSocketName(
	provideTopStackName fx.Option,
	provideWebSocketStackName fx.Option,
	provideBottomStackStackName fx.Option,
) fx.Option {
	fxOptions := []fx.Option{
		ProvideStackName(WebSocketName),
		fx.Provide(
			fx.Annotated{
				Group: "TransportFactory",
				Target: NewTransportFactory(
					WebSocketName,
					TopStackName,
					WebSocketStackName,
					BottomStackStackName,
				),
			},
		),
	}
	if provideTopStackName != nil {
		fxOptions = append(fxOptions, provideTopStackName)
	}
	if provideWebSocketStackName != nil {
		fxOptions = append(fxOptions, provideWebSocketStackName)
	}
	if provideBottomStackStackName != nil {
		fxOptions = append(fxOptions, provideBottomStackStackName)
	}
	return fx.Options(fxOptions...)
}
