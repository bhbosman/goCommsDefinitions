package goCommsDefinitions

import "go.uber.org/fx"

func ProvideTransportFactoryForOnlySSHStack(
	//provideTopStackName fx.Option,
	provideSshStackName fx.Option,
	//provideBottomStackStackName fx.Option,
) fx.Option {
	fxOptions := []fx.Option{
		ProvideStackName(TransportFactoryOnlySSHStack),
		fx.Provide(
			fx.Annotated{
				Group: "TransportFactory",
				Target: NewTransportFactory(
					TransportFactoryOnlySSHStack,
					//TopStackName,
					SshStackName,
					//BottomStackStackName,
				),
			}),
	}
	//if provideTopStackName != nil {
	//	fxOptions = append(fxOptions, provideTopStackName)
	//}
	if provideSshStackName != nil {
		fxOptions = append(fxOptions, provideSshStackName)
	}
	//if provideBottomStackStackName != nil {
	//	fxOptions = append(fxOptions, provideBottomStackStackName)
	//}
	return fx.Options(fxOptions...)
}
