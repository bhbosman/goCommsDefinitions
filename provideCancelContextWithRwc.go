package goCommsDefinitions

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"io"
)

func ProvideCancelContextWithRwc(cancelContext context.Context) fx.Option {
	return fx.Provide(
		fx.Annotated{
			Target: func(
				params struct {
					fx.In
					Lifecycle fx.Lifecycle
					Logger    *zap.Logger
					Rwc       io.ReadWriteCloser
				},
			) (context.Context, context.CancelFunc, ICancellationContext, error) {
				ctx, cancelFunc := context.WithCancel(cancelContext)
				cancellationContextInstance := NewCancellationContext(cancelFunc, ctx, params.Logger, params.Rwc)
				params.Lifecycle.Append(
					fx.Hook{
						OnStart: nil,
						OnStop: func(ctx context.Context) error {
							cancellationContextInstance.Cancel()
							return nil
						},
					},
				)
				return ctx,
					cancellationContextInstance.Cancel,
					cancellationContextInstance,
					nil
			},
		},
	)
}
