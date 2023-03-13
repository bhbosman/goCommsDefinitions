package goCommsDefinitions

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
)

func InvokeListenerClose() fx.Option {
	return fx.Invoke(
		func(
			params struct {
				fx.In
				NetListener net.Listener
				Lifecycle   fx.Lifecycle
			},
		) {
			params.Lifecycle.Append(
				fx.Hook{
					OnStart: nil,
					OnStop: func(ctx context.Context) error {
						err := params.NetListener.Close()
						return err
					},
				},
			)
		},
	)
}

func InvokeCancelContext() fx.Option {
	return fx.Invoke(
		func(
			params struct {
				fx.In
				Lifecycle           fx.Lifecycle
				CancellationContext ICancellationContext
			},
		) error {
			params.Lifecycle.Append(
				fx.Hook{
					OnStart: nil,
					OnStop: func(ctx context.Context) error {
						params.CancellationContext.Cancel()
						return nil
					},
				},
			)
			return nil
		},
	)
}
func ProvideCancelContext(cancelContext context.Context) fx.Option {
	return fx.Provide(
		fx.Annotated{
			Target: func(
				params struct {
					fx.In
					Logger         *zap.Logger
					ConnectionName string `name:"ConnectionName"`
				},
			) (context.Context, context.CancelFunc, ICancellationContext, error) {
				ctx, cancelFunc := context.WithCancel(cancelContext)
				cancelInstance := NewCancellationContext(params.ConnectionName, cancelFunc, ctx, params.Logger, nil)
				return ctx, cancelInstance.Cancel, cancelInstance, nil
			},
		},
	)
}
