package goCommsDefinitions

import "go.uber.org/fx"

func ProvideTransportFactoryForUnCompressedName(
	provideTopStackName fx.Option,
	providePingPongStackName fx.Option,
	provideProtoBufferStack fx.Option,
	provideMessageNumberStackName fx.Option,
	provideMessageBreakerStackName fx.Option,
	provideBottomStackStackName fx.Option,
) fx.Option {
	fxOptions := []fx.Option{
		ProvideStackName(TransportFactoryUnCompressedName),
		fx.Provide(
			fx.Annotated{
				Group: "TransportFactory",
				Target: NewTransportFactory(
					TransportFactoryUnCompressedName,
					TopStackName,
					PingPongStackName,
					ProtoBufferStackName,
					MessageNumberStackName,
					MessageBreakerStackName,
					BottomStackStackName,
				),
			},
		),
	}
	if provideTopStackName != nil {
		fxOptions = append(fxOptions, provideTopStackName)
	}
	if providePingPongStackName != nil {
		fxOptions = append(fxOptions, providePingPongStackName)
	}
	if provideProtoBufferStack != nil {
		fxOptions = append(fxOptions, provideProtoBufferStack)
	}
	if provideMessageNumberStackName != nil {
		fxOptions = append(fxOptions, provideMessageNumberStackName)
	}
	if provideMessageBreakerStackName != nil {
		fxOptions = append(fxOptions, provideMessageBreakerStackName)
	}
	if provideBottomStackStackName != nil {
		fxOptions = append(fxOptions, provideBottomStackStackName)
	}

	return fx.Options(fxOptions...)
}
