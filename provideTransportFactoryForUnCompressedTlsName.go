package goCommsDefinitions

import "go.uber.org/fx"

func ProvideTransportFactoryForUnCompressedTlsName(
	provideTopStackName fx.Option,
	providePingPongStackName fx.Option,
	provideProtoBufferStack fx.Option,
	provideMessageNumberStackName fx.Option,
	provideMessageBreakerStackName fx.Option,
	provideTlsStackName fx.Option,
	provideBottomStackStackName fx.Option,
) fx.Option {
	fxOptions := []fx.Option{
		ProvideStackName(TransportFactoryUnCompressedTlsName),
		fx.Provide(fx.Annotated{
			Group: "TransportFactory",
			Target: NewTransportFactory(
				TransportFactoryUnCompressedTlsName,
				TopStackName,
				PingPongStackName,
				ProtoBufferStackName,
				MessageNumberStackName,
				MessageBreakerStackName,
				TlsStackName,
				BottomStackStackName),
		}),
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
	if provideTlsStackName != nil {
		fxOptions = append(fxOptions, provideTlsStackName)
	}
	if provideBottomStackStackName != nil {
		fxOptions = append(fxOptions, provideBottomStackStackName)
	}
	return fx.Options(fxOptions...)
}
