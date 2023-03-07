package goCommsDefinitions

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideCancelContext(cancelContext context.Context) fx.Option {
	return fx.Provide(
		fx.Annotated{
			Target: func(
				params struct {
					fx.In
					Lifecycle fx.Lifecycle
					Logger    *zap.Logger
				},
			) (context.Context, context.CancelFunc, ICancellationContext, error) {
				ctx, cancelFunc := context.WithCancel(cancelContext)
				cancellationContext := newCancellationContext(cancelFunc, ctx, params.Logger, nil)
				params.Lifecycle.Append(
					fx.Hook{
						OnStart: nil,
						OnStop: func(ctx context.Context) error {
							cancellationContext.Cancel()
							return nil
						},
					},
				)
				return ctx,
					func() {
						cancellationContext.Cancel()
					},
					cancellationContext,
					nil
			},
		},
	)
}
