package goCommsDefinitions

import "go.uber.org/fx"

func ProvideTransportFactoryForSshChannelSession(
	topStackProvider fx.Option,
	//sshChannelSessionProvider fx.Option,
	bottomStackProvider fx.Option,
) fx.Option {
	var fxOption = []fx.Option{
		ProvideStackName(TransportFactoryForSshChannelSession),
		fx.Provide(
			fx.Annotated{
				Group: "TransportFactory",
				Target: NewTransportFactory(
					TransportFactoryForSshChannelSession,
					TopStackName,
					//SshChannelSession,
					BottomStackStackName),
			},
		),
	}
	if topStackProvider != nil {
		fxOption = append(fxOption, topStackProvider)
	}
	//if sshChannelSessionProvider != nil {
	//	fxOption = append(fxOption, sshChannelSessionProvider)
	//}
	if bottomStackProvider != nil {
		fxOption = append(fxOption, bottomStackProvider)
	}
	return fx.Options(fxOption...)
}

func ProvideTransportFactoryForEmptyName(
	topStackProvider fx.Option,
	bottomStackProvider fx.Option,
) fx.Option {
	var fxOption = []fx.Option{
		fx.Provide(
			fx.Annotated{
				Group: "TransportFactory",
				Target: NewTransportFactory(
					TransportFactoryEmptyName,
					TopStackName,
					BottomStackStackName),
			},
		),
	}
	if topStackProvider != nil {
		fxOption = append(fxOption, topStackProvider)
	}
	if bottomStackProvider != nil {
		fxOption = append(fxOption, bottomStackProvider)
	}
	return fx.Options(fxOption...)
}
