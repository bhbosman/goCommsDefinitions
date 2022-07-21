package goCommsDefinitions

import "go.uber.org/fx"

func ProvideTransportFactoryForCompressedTlsName(
	provideTopStackName fx.Option,
	providePingPongStackName fx.Option,
	provideProtoBufferStack fx.Option,
	provideCompressionStackName fx.Option,
	provideMessageNumberStackName fx.Option,
	provideMessageBreakerStackName fx.Option,
	provideTlsStackName fx.Option,
	provideBottomStackStackName fx.Option,
) fx.Option {
	fxOptions := []fx.Option{
		ProvideStackName(TransportFactoryCompressedTlsName),
		fx.Provide(
			fx.Annotated{
				Group: "TransportFactory",
				Target: NewTransportFactory(
					TransportFactoryCompressedTlsName,
					TopStackName,
					PingPongStackName,
					ProtoBufferStackName,
					CompressionStackName,
					MessageNumberStackName,
					MessageBreakerStackName,
					TlsStackName,
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
	if provideCompressionStackName != nil {
		fxOptions = append(fxOptions, provideCompressionStackName)
	}
	if provideMessageNumberStackName != nil {
		fxOptions = append(fxOptions, provideMessageNumberStackName)
	}
	if provideMessageBreakerStackName != nil {
		fxOptions = append(fxOptions, provideMessageBreakerStackName)
	}
	if provideTlsStackName != nil {
		fxOptions = append(fxOptions, provideTlsStackName)
	}
	if provideBottomStackStackName != nil {
		fxOptions = append(fxOptions, provideBottomStackStackName)
	}

	return fx.Options(fxOptions...)
}
